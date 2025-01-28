# connect4

Connect4 in the terminal, written in Go.

![screenshot](images/screenshot.png)

## Usage

The "game" module contains the bare-bones of a connect4 game. You can run a demo with:

```sh
go run cmd/game/main.go
```

## Proto

Experimenting with gRPC and protocol buffers for multiplayer.

```sh
protoc --go_out=generated --go-grpc_out=generated proto/connect4.proto
```

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```
