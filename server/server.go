package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/jasonsalas/grpc-calc/calculator"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "the server port")
)

// server satisfies the calculator.CalculatorServer interface
type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("sum: %v + %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() + in.GetOperand2()

	return &pb.AddResponse{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	log.Printf("difference: %v - %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() - in.GetOperand2()

	return &pb.SubtractResponse{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, in *pb.MultplyRequest) (*pb.MultiplyResponse, error) {
	log.Printf("product: %v * %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() * in.GetOperand2()

	return &pb.MultiplyResponse{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.DivisionRequest) (*pb.DivisionResponse, error) {
	log.Printf("quotient: %v / %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() / in.GetOperand2()

	return &pb.DivisionResponse{Result: result}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	pb.RegisterCalculatorServer(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
