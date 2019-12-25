package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/muchiko/go_grpc_chat/pb"
	"google.golang.org/grpc"
)

func Run() {
	conn, err := grpc.Dial("127.0.0.1:50080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	ctx := context.Background()
	stream, err := client.Transport(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go receive(stream)

	err = send(stream)
	if err != nil {
		log.Fatal(err)
	}
}

func send(stream pb.ChatService_TransportClient) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			text := scanner.Text()

			rep := regexp.MustCompile(`\s*/\s*`)
			result := rep.Split(text, -1)

			cmd := result[0]
			room_id := result[1]
			message := result[2]

			err := stream.Send(&pb.Request{
				Cmd:     cmd,
				RoomId:  room_id,
				Message: message,
			})
			if err != nil {
				stream.CloseSend()
				return err
			}
		}
	}
}

func receive(stream pb.ChatService_TransportClient) {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		fmt.Println(resp.Message)
	}
}
