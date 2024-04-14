FROM golang:1.20

WORKDIR /wallet-core-event-listener

RUN apt-get update && apt-get install -y librdkafka-dev

# Copy the local package files to the container's workspace.
ADD . .

# Build the application
RUN go build -o listener-main cmd/walletcorelistener/main.go