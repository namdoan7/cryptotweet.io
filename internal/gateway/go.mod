module github.com/levinhne/cryptotweet.io/internal/gateway

go 1.18

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0
	github.com/levinhne/cryptotweet.io/internal/common v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/k0kubun/pp/v3 v3.1.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220310020820-b874c991c1a5 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220317150908-0efb43f6373e // indirect
)

replace github.com/levinhne/cryptotweet.io/internal/common => ../common
