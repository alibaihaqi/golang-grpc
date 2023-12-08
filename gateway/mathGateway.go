package gateway

import (
	"context"
	mv1 "github.com/alibaihaqi/golang-grpc/proto/math/v1"
	"io"
)

type MathGateway struct {
	mv1.UnimplementedMathServiceServer
}

func (m *MathGateway) Add(_ context.Context, req *mv1.AddRequest) (*mv1.AddResponse, error) {
	return &mv1.AddResponse{
		Result: req.GetFirstNumber() + req.GetSecondNumber(),
	}, nil
}

func (m *MathGateway) TotalNumber(stream mv1.MathService_TotalNumberServer) error {
	var total int32 = 0

	for {
		data, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&mv1.TotalNumberResponse{
				ResultNumber: total,
			})
		}

		if err != nil {
			return err
		}

		total += data.Number
	}
}
