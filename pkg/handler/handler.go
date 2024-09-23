package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/qinyul/go-grpc-demo/pkg/model"
	pb "github.com/qinyul/go-grpc-demo/pkg/service/proto"
)

type Handler struct {
	mux               *http.ServeMux
	grpcServiceClient pb.ItemServiceClient
}

func NewHandler(c pb.ItemServiceClient) *Handler {
	h := &Handler{
		mux:               http.NewServeMux(),
		grpcServiceClient: c,
	}
	h.routes()
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) routes() {
	h.mux.HandleFunc("/", h.handleRoot())
}

func (h *Handler) handleRootGET(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	items, err := h.grpcServiceClient.GetItems(context.Background(), &pb.Empty{})

	if items == nil {
		fmt.Println("handleRootGET:: item data empty")
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Item data not found",
		})
	} else {
		if err != nil {
			log.Printf("handleRootGET:: Error - %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error",
			})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(items)
		}
	}

}

func (h *Handler) handleRootPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req model.Item

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Println("handleRootPost:: failed to encode request")

		json.NewEncoder(w).Encode(model.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	} else {
		request := pb.ItemRequest{
			Name: req.Name,
		}
		item, err := h.grpcServiceClient.CreateUser(context.Background(), &request)

		if item == nil || err != nil {
			fmt.Println("handleRootPost:: failed to create data")
			json.NewEncoder(w).Encode(model.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error",
			})
		} else {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
		}
	}
}

func handleUnsportedMethod(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Unsported Method")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(model.ErrorResponse{
		Status:  http.StatusBadRequest,
		Message: "Unsported Method",
	})
}
func (h *Handler) handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			h.handleRootGET(w)
		case "POST":
			h.handleRootPost(w, r)
		default:
			handleUnsportedMethod(w)
		}

	}
}
