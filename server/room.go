package server

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/muchiko/go_grpc_chat/pb"
)

type Room struct {
	room_id string
	streams map[*pb.ChatService_TransportServer]bool
}

type room_regist struct {
	mqdb    *mqtt.Client
	room_id string
	stream  *pb.ChatService_TransportServer
}

type room_publish struct {
	mqdb    *mqtt.Client
	room_id string
	stream  *pb.ChatService_TransportServer
	message string
}

func createRoom(room_id string) *Room {
	return &Room{
		room_id: room_id,
		streams: make(map[*pb.ChatService_TransportServer]bool),
	}
}

func (r *Room) appendStream(stream *pb.ChatService_TransportServer) {
	r.streams[stream] = true
}

func (r *Room) deleteStream(stream *pb.ChatService_TransportServer) {
	delete(r.streams, stream)
}
