# Go RabbitMQ JSON Messaging Example

This project demonstrates a simple client-server setup in Go that communicates using JSON messages over RabbitMQ.

## Project Structure

```
rabbitmq-json-example/
├── cmd/
│   ├── client/
│   │   └── main.go       # Client application to send messages
│   └── server/
│       └── main.go       # Server application to receive messages
├── pkg/
│   ├── messaging/
│   │   └── common.go     # Shared constants (RabbitMQ URL, queue name)
│   └── types/
│       └── message.go    # Shared struct for the JSON message
├── docker-compose.yml    # Docker Compose to run RabbitMQ
├── go.mod
└── go.sum
```

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.21 or later)
- [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/)

## How to Run

### 1. Start RabbitMQ

In the project root directory, start the RabbitMQ container using Docker Compose.

```sh
docker-compose up -d
```

This will start a RabbitMQ instance with the management UI available at `http://localhost:15672`. You can log in with `guest`/`guest`.

### 2. Run the Server

Open a new terminal window, navigate to the project root, and run the server:

```sh
go run ./cmd/server/
```

The server will start and wait for messages.

### 3. Run the Client

Open a third terminal window. You can send a default message or provide your own as a command-line argument.

```sh
# Send a default message
go run ./cmd/client/

# Send a custom message
go run ./cmd/client/ "This is a custom message from the client"
```

You will see the server's terminal log the message it received.

### 4. Clean Up

To stop and remove the RabbitMQ container:

```sh
docker-compose down
```

