# ♟️ chess websocket server

This project represents a simple websocket server to play chess.

- [Local setup](#local-setup)
- [Interacting with websocket server to play chess](#interacting-with-websocket-server-to-play-chess)
- [Chess moves](#chess-moves)
- [Features](#features)
- [Limitations](#limitations)
- [References](#references)

## Local setup

There are two major ways to locally setup the project to have it up and running:

- [Build from source](#build-from-source)
- [Build with Docker](#build-from-docker)

### Build from source

Requirements:

- `golang 1.19`
- `PORT` environment exported

Export `PORT` environmental variable. The port can be any port you would like.

```
export PORT=8080
```

Install dependencies:

```bash
go mod download
```

Run the project:

```bash
go run .
```

### Build with Docker

Requirements:

- `docker`

Build the docker image:

```bash
docker build -t chess .
```

Run the project with `PORT` env set:

```bash
docker run -d -e PORT=8080 -p 8080:8080 chess
```

### Interacting with websocket server to play chess

> a tool to interact with websockets in your terminal, such as [`websocat`](https://github.com/vi/websocat), might be handy to test the server

To play chess, players need to be matched with other player.

Firstly, connect using websocket to the endpoint `ws://localhost:8080/rooms`.

You can use `websocat` or any other CLI tool you like to do so if you are running the server locally.

```bash
websocat ws://localhost:8080/rooms
```

After two successful connections, the clients (e.g. player) will receive a response from the server.
The response will be of the following type:

```
9c954450-ad7b-4dcc-ab2f-6c556c0835ef
```

This is the UUID of the game session.

Secondly, when UUID is received connect to the next endpoint - `ws://localhost:8080/rooms/9c954450-ad7b-4dcc-ab2f-6c556c0835ef`.
The UUID should be placed after `/rooms/` path.

After the successful connection of 2 clients, the client will receive its own chess color, either `white` or `black`.

From now on, players can exchange moves to play chess.

> If only one user connects, then the server will not send any color information.
> That's because the server waits for second player to join. After both players join,
> the players will receive own colors.

## Chess moves

The players exchange with text messages to indicate their chess move.
As of chess move notation - [Long algebraic notation](<https://en.wikipedia.org/wiki/Algebraic_notation_(chess)#Long_algebraic_notation>) is used.

Players need to send chess move messages strictly by the defined notion. If the chess engine fails
to identify the chess move, then the player will receive an error message - the user will be prompted
to send a valid move.

## Features

- Supports many concurrent games.

## Limitations

- No player reconnect mechanism. If one of user's connection interrupts, then the game session will end.

## References

- Chess Engine - [notnil/chess](https://github.com/notnil/chess).
- Websocket library - [gorilla/websocket](https://github.com/gorilla/websocket).
