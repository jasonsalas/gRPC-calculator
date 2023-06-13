package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/jasonsalas/grpc-calc/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr          = flag.String("addr", "localhost:50052", "the address to connect to")
	operand1      = flag.Int("op1", 2, "the first operand")
	operand2      = flag.Int("op2", 4, "the second operand")
)

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
	result, err := client.Add(ctx, &pb.AddRequest{Operand1: int32(*operand1), Operand2: int32(*operand2)})
	if err != nil {
		log.Fatalf("unable to add integers: %v", err)
	}
	log.Printf("calling Add(): %v", result.GetResult())
}
