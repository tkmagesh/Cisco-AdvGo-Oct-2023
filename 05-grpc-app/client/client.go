package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	proto "github.com/tkmagesh/cisco-advgo-oct-2023/05-grpc-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	options := grpc.WithTransportCredentials(insecure.NewCredentials())
	clientConn, err := grpc.Dial("localhost:50051", options)
	if err != nil {
		log.Fatalln(err)
	}

	appServiceClient := proto.NewAppServiceClient(clientConn)
	ctx := context.Background()

	// doRequestResponse(ctx, appServiceClient)
	// doServerStreaming(ctx, appServiceClient)
	doClientStreaming(ctx, appServiceClient)
}

func doRequestResponse(ctx context.Context, appServiceClient proto.AppServiceClient) {
	addRequest := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	addResponse, err := appServiceClient.Add(ctx, addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Add Result : ", addResponse.GetResult())
}

func doServerStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {

	// create the request object
	primeRequest := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}

	// client stream object to receive the stream of responses
	var clientStream proto.AppService_GeneratePrimesClient
	clientStream, err := appServiceClient.GeneratePrimes(ctx, primeRequest)
	if err != nil {
		log.Fatalln(err)
	}

	// receive the stream of responses as and when the server sends
	for {
		primeResponse, err := clientStream.Recv()
		if err == io.EOF {
			fmt.Println("All the responses received")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Prime No : ", primeResponse.GetPrimeNo())
	}

}

func doClientStreaming(ctx context.Context, appServiceClient proto.AppServiceClient) {
	data := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var clientStream proto.AppService_CalculateAverageClient
	clientStream, err := appServiceClient.CalculateAverage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Starting calculating the average")
	for _, value := range data {
		req := &proto.AverageRequest{
			No: value,
		}
		fmt.Printf("Sending no : %d\n", value)
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Finished sending the data")
	averageResponse, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Average : %d\n", averageResponse.GetResult())
}
