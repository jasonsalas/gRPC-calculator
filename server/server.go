package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "github.com/jasonsalas/grpc-calc/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port            = flag.Int("port", 50052, "the server port")
	ErrDivideByZero = errors.New("you cannot divide by zero")
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func makeAccessLog() *log.Logger {
	logPath, err := filepath.Abs("./logs/access.log")
	if err != nil {
		log.Fatalf("could not create the logfile: %v", err)
	}
	f, _ := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0775)
	return log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	accesslog := makeAccessLog()
	log.Printf("sum: %v + %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() + in.GetOperand2()
	accesslog.Printf("calling Add(): %v + %v = %v", in.GetOperand1(), in.GetOperand2(), result)

	return &pb.AddResponse{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, in *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	accesslog := makeAccessLog()
	log.Printf("difference: %v - %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() - in.GetOperand2()
	accesslog.Printf("calling Subtract(): %v - %v = %v", in.GetOperand1(), in.GetOperand2(), result)

	return &pb.SubtractResponse{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, in *pb.MultplyRequest) (*pb.MultiplyResponse, error) {
	accesslog := makeAccessLog()
	log.Printf("product: %v * %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() * in.GetOperand2()
	accesslog.Printf("calling Multiply(): %v * %v = %v", in.GetOperand1(), in.GetOperand2(), result)

	return &pb.MultiplyResponse{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, in *pb.DivisionRequest) (*pb.DivisionResponse, error) {
	if in.GetOperand2() == 0 {
		return nil, ErrDivideByZero
	}
	accesslog := makeAccessLog()
	log.Printf("quotient: %v / %v", in.GetOperand1(), in.GetOperand2())
	result := in.GetOperand1() / in.GetOperand2()
	accesslog.Printf("calling Divide(): %v / %v = %v", in.GetOperand1(), in.GetOperand2(), result)

	return &pb.DivisionResponse{Result: result}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	pb.RegisterCalculatorServer(srv, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
