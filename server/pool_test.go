package server

import (
	"testing"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/mock/gomock"
	"github.com/muchiko/go_grpc_chat/mock"
)

var p = newPool()
var room_id = "room"
var qos = byte(1)
var retained = false
var message = "message"
var room = createRoom(room_id)

func TestSubscribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	token := mock.NewMockToken(ctrl)
	token.EXPECT().Wait().Return(true)
	token.EXPECT().Error().Return(nil)

	client := mock.NewMockClient(ctrl)
	client.EXPECT().Subscribe(room_id, qos, gomock.AssignableToTypeOf(p._subscribe(room))).Return(token)

	if err := testSubscribe(client); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func testSubscribe(client mqtt.Client) error {
	p.subscribe(&client, room)
	return nil
}

func xxTestPublish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	token := mock.NewMockToken(ctrl)
	token.EXPECT().Wait().Return(true)
	token.EXPECT().Error().Return(nil)

	client := mock.NewMockClient(ctrl)
	client.EXPECT().Publish(room_id, qos, retained, message).Return(token)

	if err := testPublish(client); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func testPublish(client mqtt.Client) error {
	p._publish(&client, room_id, qos, retained, message)
	return nil
}
