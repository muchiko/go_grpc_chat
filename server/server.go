package server

import (
	"fmt"
	"io"
	"log"
	"net"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/muchiko/go_grpc_chat/pb"
	"google.golang.org/grpc"
)

func Run() {
	lis, err := net.Listen("tcp", ":50080")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://0.0.0.0:1883")
	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	p := newPool()
	go p.run()

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{
		mqdb: &client,
		pool: p,
	})
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	mqdb *mqtt.Client
	pool *Pool
}

func (s *server) Transport(stream pb.ChatService_TransportServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Println(in.Cmd)
		fmt.Println(in.RoomId)
		fmt.Println(in.Message)

		switch in.Cmd {
		case "sub":
			s.pool.register <- room_regist{mqdb: s.mqdb, room_id: in.RoomId, stream: &stream}
		case "pub":
			s.pool.publish <- room_publish{mqdb: s.mqdb, room_id: in.RoomId, stream: &stream, message: in.Message}
		default:
			return nil
		}
	}
}
