package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	proto "github.com/tkmagesh/cisco-advgo-oct-2023/05-grpc-app/proto"
	"google.golang.org/grpc"
)

type AppServiceImpl struct {
	proto.UnimplementedAppServiceServer
}

func (asi *AppServiceImpl) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	fmt.Printf("Received Add req, x = %d and y = %d\n", x, y)
	result := x + y
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func (asi *AppServiceImpl) GeneratePrimes(req *proto.PrimeRequest, serverStream proto.AppService_GeneratePrimesServer) error {
	start := req.GetStart()
	end := req.GetEnd()
	fmt.Printf("Received Prime req, start = %d and end = %d\n", start, end)
	for no := start; no <= end; no++ {
		if isPrime(no) {
			res := &proto.PrimeResponse{
				PrimeNo: no,
			}
			fmt.Println("Sending Prime No :", no)
			err := serverStream.Send(res)
			if err != nil {
				return err
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func (asi *AppServiceImpl) CalculateAverage(serverStream proto.AppService_CalculateAverageServer) error {
	var sum, count int32
	for {
		req, err := serverStream.Recv()
		if err == io.EOF {
			fmt.Println("All the data received")
			break
		}
		if err != nil {
			return err
		}
		no := req.GetNo()
		fmt.Println("Received No :", no)
		count += 1
		sum += no
	}
	avg := sum / count
	res := &proto.AverageResponse{
		Result: avg,
	}
	serverStream.SendAndClose(res)
	return nil
}

func main() {
	asi := &AppServiceImpl{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, asi)
	grpcServer.Serve(listener)
}
