FROM golang:1.20

WORKDIR /wallet-core

RUN apt-get update && apt-get install -y librdkafka-dev

# Copy the local package files to the container's workspace.
ADD . .

CMD ["tail", "-f", "/dev/null"]
## Build the application
RUN go build -o main cmd/walletcore/main.go
#
## Run the binary program produced by `go build`
CMD ["./main"]