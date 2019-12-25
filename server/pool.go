package server

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/muchiko/go_grpc_chat/pb"
)

type Pool struct {
	rooms    map[string]*Room
	register chan room_regist
	publish  chan room_publish
}

func newPool() *Pool {
	return &Pool{
		rooms:    make(map[string]*Room),
		register: make(chan room_regist),
		publish:  make(chan room_publish),
	}
}

func (p *Pool) run() {
	for {
		select {
		case rr := <-p.register:
			var room *Room
			if val, ok := p.rooms[rr.room_id]; ok {
				room = val
			} else {
				room = createRoom(rr.room_id)
				p.rooms[rr.room_id] = room
			}

			room.appendStream(rr.stream)

			err := p.subscribe(rr.mqdb, room)
			if err != nil {
				fmt.Println(err)
			}

		case rp := <-p.publish:
			topic := rp.room_id
			qos := byte(1)
			retained := false
			token := (*rp.mqdb).Publish(topic, qos, retained, rp.message)
			if token.Wait() && token.Error() != nil {
				fmt.Println(token.Error())
			}
		}
	}
}

func (p *Pool) subscribe(mqdb *mqtt.Client, room *Room) error {
	topic := room.room_id
	qos := byte(1)
	token := (*mqdb).Subscribe(topic, qos, func(client mqtt.Client, msg mqtt.Message) {
		for key, _ := range room.streams {
			(*key).Send(&pb.Payload{
				Message: string(msg.Payload()),
			})
		}
	})
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
