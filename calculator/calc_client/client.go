package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/kylezs/grpc-go-course/calculator/calcpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Running the calculator client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calcpb.NewCalcServiceClient(cc)

	// getSumFromServer(c, 2, 1332)

	// getPrimeDecompFromServer(c, 200)

	// getComputedAverageFromServer(c)

	doErrorUnary(c)

}

func getComputedAverageFromServer(c calcpb.CalcServiceClient) {
	fmt.Println("Starting client streaming RPC to average numbers")
	numbers := []int64{3, 5, 9, 54, 23}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling compute average: %v", err)
	}

	for _, n := range numbers {
		fmt.Printf("Sending n: %v\n", n)
		stream.Send(&calcpb.AverageRequest{
			Number: n,
		})
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from ComputeAverage %v", err)
	}
	fmt.Printf("The average is: %v\n", res)
}

func getPrimeDecompFromServer(c calcpb.CalcServiceClient, x int32) {
	fmt.Println("Starting to stream prime numbers...")

	req := &calcpb.PrimeDecompRequest{
		Integer: x,
	}

	resStream, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling prime decomposition: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Server closed the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while readin the stream: %v", err)
		}
		// if all is good
		log.Printf("%d", msg.GetPrime())
	}

}

func getSumFromServer(c calcpb.CalcServiceClient, x int32, y int32) {
	fmt.Println("client about to call Sum()")

	req := &calcpb.CalcRequest{
		Num1: x,
		Num2: y,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Response from Sum: %v", res.Answer)
}

func doErrorUnary(c calcpb.CalcServiceClient) {
	fmt.Println("Starting to do a unary rpc")

	doErrorCall(c, 10)

	doErrorCall(c, -2)

}

func doErrorCall(c calcpb.CalcServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calcpb.SquareRootRequest{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
			}
		} else {
			log.Fatalf("Big error calling square root: %v", err)
		}
	}
	fmt.Printf("Result of square root of %v: %v", n, res.GetNumberRoot())
}
