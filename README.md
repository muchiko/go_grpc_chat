# go_grpc_chat

```
brew install mosquitto
brew services start mosquitto
/usr/local/opt/mosquitto/sbin/mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf
```


```
mockgen github.com/muchiko/go_grpc_chat/pb ChatServiceClient,ChatService_TransportClient > ./mock_pb/mock_chat.pb.go
```
