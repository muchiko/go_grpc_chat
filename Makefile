.PHONY: fmt
fmt: 
	go fmt client/*.go
	go fmt server/*.go
	go fmt *.go

.PHONY: server
server: 
	go run server.go

.PHONY: client
client: 
	go run client.go
