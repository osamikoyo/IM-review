package server

import (
	"context"
	"net/http"

	"github.com/osamikoyo/IM-review/internal/data"
	"github.com/osamikoyo/IM-review/internal/data/models"
	"github.com/osamikoyo/IM-review/pkg/proto/pb"
)

type Server struct{
	pb.UnimplementedReviewServiceServer
	Storage *data.Storage
}

func (s *Server) Add(_ context.Context, req *pb.AddReviewReq) (*pb.Response, error){
	err := s.Storage.Add(models.ToModels(req.Review))
	if err != nil{
		return &pb.Response{
			Error: err.Error(),
			Status: http.StatusInternalServerError,
		}, err
	}

	return &pb.Response{
		Error: "",
		Status: http.StatusOK,
	}, nil
}
func (s *Server) Get(_ context.Context, req *pb.GetReviewsReq) (*pb.GetReviewsResp, error){
	reviews, err := s.Storage.Get(req.ProductID)
	if err != nil{
		return &pb.GetReviewsResp{
			Response: &pb.Response{
				Error: err.Error(),
				Status: http.StatusInternalServerError,
			},
		}, err
	}

	return &pb.GetReviewsResp{
		Reviews: models.ToPB(reviews),
		Response: &pb.Response{
			Error: "",
			Status: http.StatusOK,
		},
	}, nil
}