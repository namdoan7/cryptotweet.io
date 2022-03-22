module github.com/levinhne/cryptotweet.io/internal/gateway

go 1.17

replace github.com/levinhne/cryptotweet.io/internal/common => ../common

require (
	github.com/levinhne/cryptotweet.io/internal/common v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.8.0
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220308174144-ae0e22291548 // indirect
)
