#Dockerfile for server and client

#Use the official golang image
FROM golang:1.23 AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Install dependecies
RUN go mod download

#Copy the source code
COPY . .

#Copy the config file
COPY config.json ./config.json

# Install protobuf compiler and Go plugins
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    && go get google.golang.org/protobuf/cmd/protoc-gen-go \
    && go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


# Ensure the Go binaries are in the PATH
ENV PATH="$PATH:/go/bin"

# Generate the protobuf code
RUN protoc --go_out=. --go-grpc_out=. proto/item.proto

#Build the server
RUN go build -o ./cmd/microservice/server ./cmd/microservice/server

#Build the client
RUN go build -o ./cmd/microservice/client ./cmd/microservice/client

# Build the server with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/microservice/server ./cmd/microservice/server

# Build the client with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/microservice/client ./cmd/microservice/client

#Use a smaller image to run server and client
FROM alpine:latest

#Set the current working directory inside the container
WORKDIR /app

# Copy the compiled binaries from the builder stage
COPY --from=builder /app/cmd/microservice/server ./cmd/microservice/server
COPY --from=builder /app/cmd/microservice/client ./cmd/microservice/client 


# Expose the port the server runs on
EXPOSE 8000