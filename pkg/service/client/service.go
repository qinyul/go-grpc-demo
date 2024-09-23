package service

import (
	"fmt"

	pb "github.com/qinyul/go-grpc-demo/pkg/service/proto"
)

type Service struct {
	sc pb.ItemServiceClient
}

func (s Service) GetItems() {
	fmt.Println()
}
