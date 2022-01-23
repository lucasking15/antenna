package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "antenna/api"
)

type AntennaService struct {
	pb.UnimplementedAntennaServer

	log *log.Helper
}

func NewAntennaService(logger log.Logger) *AntennaService {
	return &AntennaService{log: log.NewHelper(logger)}
}

func (s *AntennaService) ListAntenna(ctx context.Context, req *pb.ListAntennaRequest) (*pb.ListAntennaReply, error) {
	return &pb.ListAntennaReply{Code: 12, Msg: "12"}, nil
}
