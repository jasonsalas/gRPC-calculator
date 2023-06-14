package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "github.com/jasonsalas/grpc-calc/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr     = flag.String("addr", "localhost:50052", "the address to connect to")
	operator = flag.String("op", "add", "math operator - addition, subtraction, multiplication, division")
	operand1 = flag.Int("op1", 2, "the first operand")
	operand2 = flag.Int("op2", 4, "the second operand")
)

func makeAccessLog() *log.Logger {
	logPath, err := filepath.Abs("./logs/access.log")
	if err !=nil{
		panic(err)
	}
	f, _ := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0775)
	return log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to the gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	accesslog := makeAccessLog()
	var result interface{}
	switch *operator {
	case "add":
		result, err = client.Add(ctx, &pb.AddRequest{Operand1: int32(*operand1), Operand2: int32(*operand2)})
		if err != nil {
			log.Fatalf("unable to add integers: %v", err)
		}
		log.Printf("calling Add(): %v", result.(*pb.AddResponse).GetResult())
		accesslog.Printf("calling Add(): %v", result.(*pb.AddResponse).GetResult())
	case "sub":
		result, err = client.Subtract(ctx, &pb.SubtractRequest{Operand1: int32(*operand1), Operand2: int32(*operand2)})
		if err != nil {
			log.Fatalf("unable to subtract integers: %v", err)
		}
		log.Printf("calling Sub(): %v", result.(*pb.SubtractResponse).GetResult())
		accesslog.Printf("calling Sub(): %v", result.(*pb.SubtractResponse).GetResult())
	case "mul":
		result, err = client.Multiply(ctx, &pb.MultplyRequest{Operand1: int32(*operand1), Operand2: int32(*operand2)})
		if err != nil {
			log.Fatalf("unable to multiply integers:%v", err)
		}
		log.Printf("calling Multiply(): %v", result.(*pb.MultiplyResponse).GetResult())
		accesslog.Printf("calling Multiply(): %v", result.(*pb.MultiplyResponse).GetResult())
	case "div":
		result, err = client.Divide(ctx, &pb.DivisionRequest{Operand1: int32(*operand1), Operand2: int32(*operand2)})
		if err != nil {
			log.Fatalf("unable to divide integers: %v", err)
		}
		log.Printf("calling Divide(): %v", result.(*pb.DivisionResponse).GetResult())
		accesslog.Printf("calling Divide(): %v", result.(*pb.DivisionResponse).GetResult())
	default:
		result, err = client.Add(ctx, &pb.AddRequest{Operand1: int32(*operand1), Operand2: int32(*operand2)})
		if err != nil {
			log.Fatalf("unable to add integers: %v", err)
		}
		log.Printf("calling Add(): %v", result.(*pb.AddResponse).GetResult())
		accesslog.Printf("calling Add(): %v", result.(*pb.AddResponse).GetResult())
	}
}
