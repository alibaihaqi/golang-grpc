package gateway

import (
	"context"
	mv1 "github.com/alibaihaqi/golang-grpc/proto/math/v1"
	"io"
	"math"
	"time"
)

type MathGateway struct {
	mv1.UnimplementedMathServiceServer
}

func (m *MathGateway) Add(_ context.Context, req *mv1.AddRequest) (*mv1.AddResponse, error) {
	return &mv1.AddResponse{
		Result: req.GetFirstNumber() + req.GetSecondNumber(),
	}, nil
}

func (m *MathGateway) SinCos(req *mv1.SinCosRequest, stream mv1.MathService_SinCosServer) error {
	commonDegreesList := []int{0, 30, 45, 60, 90, 120, 135, 150, 180, 210, 225, 240, 270, 300, 315, 330, 360}

	for _, degrees := range commonDegreesList {
		radians := float64(degrees) * (math.Pi / 180)

		var value float64 = 0

		if req.Method == mv1.SineCosineEnum_SINE {
			value = math.Sin(radians)
		} else if req.Method == mv1.SineCosineEnum_COSINE {
			value = math.Cos(radians)
		}

		threshold := 1e-10
		if value > 0 && value < threshold {
			value = math.Floor(value)
		} else if value < 0 && value > -threshold {
			value = math.Ceil(value)
		}

		res := &mv1.SinCosResponse{
			Degree: int32(degrees),
			Value:  float32(value),
		}

		time.Sleep(time.Second)

		if err := stream.Send(res); err != nil {
			return err
		}
	}

	return nil
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
