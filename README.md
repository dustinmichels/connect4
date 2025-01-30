# connect4

Connect4 in the terminal, written in Go.

![screenshot](images/screenshot.png)

## Usage

```sh
go run .
```

## Ambition

The goal is to add a bot for single-player, and online multiplayer node.

### Single-player

The bot will use monte carlo simulations to determine the best move.

### Multiplayer

One user must run a gRPC server. Other players can connect as clients. When you launch a server, a google sheet will be updated with an encoded IP address, so other player can find you.

Clients can be in the terminal, or a web app.

#### Proto

Experimenting with gRPC and protocol buffers for multiplayer.

To update the proto files, run:

```sh
./dev/proto.sh
```

Maybe even decentralized??

- [libp2p](https://docs.libp2p.io/guides/getting-started/go/)

[Google sheet](https://script.google.com/macros/s/AKfycbxHDhTy_UIjL51FrT6E9dKqMb1rYmdy2ZnLrRpubTdhXMgdy-fCKeKY1eSvPJuw_0s/exec)
