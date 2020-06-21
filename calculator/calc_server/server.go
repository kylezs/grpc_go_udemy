package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/reflection"

	"github.com/kylezs/grpc-go-course/calculator/calcpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	fmt.Printf("Sum is being called with %v \n", req)
	x := req.GetNum1()
	y := req.GetNum2()
	answer := x + y
	response := &calcpb.CalcResponse{
		Answer: answer,
	}

	return response, nil
}

func (*server) PrimeDecomposition(req *calcpb.PrimeDecompRequest, stream calcpb.CalcService_PrimeDecompositionServer) error {
	fmt.Printf("Prime decomposition function called with: %v\n", req)
	number := req.GetInteger()
	intNum := int(number)
	k := 2
	for intNum > 1 {
		if intNum%k == 0 {
			// Send k to the client
			res := &calcpb.PrimeDecompResponse{
				Prime: int64(k),
			}
			fmt.Printf("Sending %d\n", k)
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			intNum = intNum / k
		} else {
			k++
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calcpb.CalcService_ComputeAverageServer) error {
	fmt.Println("Compute average function was invoked with a streaming request")
	var sum int64 = 0
	var n int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			average := float32(sum) / float32(n)
			return stream.SendAndClose(&calcpb.AverageResponse{
				Average: average,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		number := req.GetNumber()
		sum += number
		n++
	}
}

func (*server) SquareRoot(ctx context.Context, req *calcpb.SquareRootRequest) (*calcpb.SquareRootResponse, error) {
	fmt.Println("Received square root RPC")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &calcpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("Running the calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterCalcServiceServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
