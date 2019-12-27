.PHONY: fmt
fmt: 
	go fmt client/*.go
	go fmt server/*.go
	go fmt *.go

.PHONY: test
test: 
	go test client/*.go -v
	go test server/*.go -v

.PHONY: server
server: 
	go run server.go

.PHONY: client
client: 
	go run client.go

.PHONY: mockgen
mockgen: 
	mockgen -package mock github.com/eclipse/paho.mqtt.golang Client,Token > ./mock/mock_mqtt.go
	mockgen -package mock github.com/muchiko/go_grpc_chat/pb ChatServiceClient,ChatService_TransportClient > ./mock/mock_client.pb.go
