package models

import "github.com/osamikoyo/IM-review/pkg/proto/pb"

type Review struct{
	ProductID uint64
	Mark uint8
	Content string
}

func ToModels(r *pb.Review) *Review {
	return &Review{
		ProductID: r.ProductID,
		Mark: uint8(r.Mark),
		Content: r.Content,
	}
}

func ToPB(r []Review) []*pb.Review {
	var res []*pb.Review
	for _, rs := range r{
		res = append(res, &pb.Review{
			ProductID: rs.ProductID,
			Content: rs.Content,
			Mark: uint32(rs.Mark),
		})
	}
}