package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/qinyul/go-grpc-demo/pkg/config"
	"github.com/qinyul/go-grpc-demo/pkg/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/qinyul/go-grpc-demo/pkg/service/proto"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	conn, connErr := grpc.NewClient("server:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if connErr != nil {
		fmt.Printf("did not connect to: %v", err)
	}

	defer conn.Close()

	fmt.Println("Connected to grpc server")

	c := pb.NewItemServiceClient(conn)

	h := handler.NewHandler(c)

	serverAddr := cfg.ServerAddress
	log.Printf("Starting server on %s", serverAddr)
	if err := http.ListenAndServe(":8000", h); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
