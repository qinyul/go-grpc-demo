package main

import (
	"fmt"
	"log"
	"net"

	"github.com/qinyul/go-grpc-demo/pkg/db"
	pb "github.com/qinyul/go-grpc-demo/pkg/service/proto"
	ss "github.com/qinyul/go-grpc-demo/pkg/service/server"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedItemServiceServer
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	db.InitDB()
	s := grpc.NewServer()

	pb.RegisterItemServiceServer(s, &ss.Service{})
	fmt.Println("Item Server is running on port ", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
