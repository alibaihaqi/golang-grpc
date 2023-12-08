package main

import (
	"fmt"
	"github.com/alibaihaqi/golang-grpc/gateway"
	mv1 "github.com/alibaihaqi/golang-grpc/proto/math/v1"
	"google.golang.org/grpc"
	"net"
)

func main() {
	p := "3000" // port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", p))
	if err != nil {
		return
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)

	mv1.RegisterMathServiceServer(s, &gateway.MathGateway{})
	err = s.Serve(listener)

	if err != nil {
		panic(err)
	}
}
