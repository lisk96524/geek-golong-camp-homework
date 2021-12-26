package service

import (
	"context"

	pb "week4/api/week4"
)

type Week4Service struct {
	pb.UnimplementedWeek4Server
}

func NewWeek4Service() *Week4Service {
	return &Week4Service{}
}

func (s *Week4Service) CreateWeek4(ctx context.Context, req *pb.CreateWeek4Request) (*pb.CreateWeek4Reply, error) {
	return &pb.CreateWeek4Reply{}, nil
}
func (s *Week4Service) UpdateWeek4(ctx context.Context, req *pb.UpdateWeek4Request) (*pb.UpdateWeek4Reply, error) {
	return &pb.UpdateWeek4Reply{}, nil
}
func (s *Week4Service) DeleteWeek4(ctx context.Context, req *pb.DeleteWeek4Request) (*pb.DeleteWeek4Reply, error) {
	return &pb.DeleteWeek4Reply{}, nil
}
func (s *Week4Service) GetWeek4(ctx context.Context, req *pb.GetWeek4Request) (*pb.GetWeek4Reply, error) {
	return &pb.GetWeek4Reply{}, nil
}
func (s *Week4Service) ListWeek4(ctx context.Context, req *pb.ListWeek4Request) (*pb.ListWeek4Reply, error) {
	return &pb.ListWeek4Reply{}, nil
}
