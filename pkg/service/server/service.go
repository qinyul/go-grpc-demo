package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/qinyul/go-grpc-demo/internal/utils"
	"github.com/qinyul/go-grpc-demo/pkg/model"
	pb "github.com/qinyul/go-grpc-demo/pkg/service/proto"
)

type Service struct {
	pb.UnimplementedItemServiceServer
}

func (s *Service) CreateUser(context context.Context, req *pb.ItemRequest) (*pb.ItemResponse, error) {
	fmt.Println("CreateItem:: Incoming request create item")
	if req.Name == "" {
		return nil, fmt.Errorf("CreateItem:: item id and name cannot be empty")
	}

	items, err := s.GetItems(context, &pb.Empty{})

	if err != nil {
		fmt.Println()
		return nil, err
	}

	res := &pb.ItemResponse{
		Id:        uuid.NewString(),
		Name:      req.Name,
		CreatedAt: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	items.Items = append(items.Items, res)

	dbPath := "data/item.json"
	utils.WriteJSONFile(dbPath, items.Items)

	return res, nil
}

func (s *Service) GetItems(ctx context.Context, req *pb.Empty) (*pb.ItemsResponse, error) {
	fmt.Println("GetItem:: Incoming request get item")
	file, err := os.Open("data/item.json")

	if err != nil {
		fmt.Println("GetItem Err:: ", err)
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatalf("GetItem:: Error reading file: %v \n", err)
	}

	var itemsCol []model.Item

	fmt.Println("GetItem:: Starting unmarshalling data")
	if err := json.Unmarshal(data, &itemsCol); err != nil {
		fmt.Printf("GetItem:: error unmarshaling data %v\n", err)
		return nil, err
	}

	itemMap := make(map[string]*pb.ItemResponse)
	fmt.Println("GetItem:: constructing item map")

	for _, item := range itemsCol {
		itemMap[item.ID] = &pb.ItemResponse{
			Id:        item.ID,
			Name:      item.Name,
			CreatedAt: item.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
		if item.UpdatedAt != nil {
			itemMap[item.ID].UpdatedAt = item.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		}
	}

	var results []*pb.ItemResponse
	for _, item := range itemMap {
		results = append(results, item)
	}

	return &pb.ItemsResponse{Items: results}, nil
}
