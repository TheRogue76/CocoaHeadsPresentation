package main

import (
	"context"
	pb "github.com/TheRogue76/CocoaHeadsPresentation/github.com/TheRogue76/CocoaHeadsPresentation/FoodCompanyGateWay"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedFoodCompanyGateWayServiceServer
}

func (s *server) GetOrderStatus(_ context.Context, request *pb.GetOrderStatusRequest) (*pb.GetOrderStatusResponse, error) {
	if request.Id == "1" {
		return &pb.GetOrderStatusResponse{Status: pb.GetOrderStatusResponse_Done}, nil
	}
	if request.Id == "2" {
		owner := "Parsa"
		return &pb.GetOrderStatusResponse{Status: pb.GetOrderStatusResponse_Pending, Owner: &owner}, nil
	}
	return &pb.GetOrderStatusResponse{Status: pb.GetOrderStatusResponse_Cancelled}, nil
}

func main() {
	list, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFoodCompanyGateWayServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", list.Addr())
	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
