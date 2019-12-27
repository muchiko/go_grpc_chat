package client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
	"github.com/muchiko/go_grpc_chat/mock"
	"github.com/muchiko/go_grpc_chat/pb"
)

var req = &pb.Request{
	Cmd:     "cmd",
	RoomId:  "room_id",
	Message: "message",
}

var msg = &pb.Payload{
	Message: "parcel",
}

func TestTransport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	stream := mock.NewMockChatService_TransportClient(ctrl)
	stream.EXPECT().Send(
		gomock.Any(),
	).Return(nil)
	stream.EXPECT().Recv().Return(msg, nil)
	stream.EXPECT().CloseSend().Return(nil)

	client := mock.NewMockChatServiceClient(ctrl)
	client.EXPECT().Transport(
		gomock.Any(),
	).Return(stream, nil)

	if err := testTransport(client); err != nil {
		t.Fatalf("Test failed: %v", err)
	}
}

func testTransport(client pb.ChatServiceClient) error {
	//	return nil
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Transport(ctx)
	if err != nil {
		return err
	}
	if err := stream.Send(req); err != nil {
		return err
	}
	if err := stream.CloseSend(); err != nil {
		return err
	}
	got, err := stream.Recv()
	if err != nil {
		return err
	}

	if !proto.Equal(got, msg) {
		return fmt.Errorf("stream.Recv() = %v, want %v", got, msg)
	}

	return nil
}
