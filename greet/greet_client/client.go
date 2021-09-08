package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ikalkali/belajar-grpc/greet/greetpb"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println("Hello im a client!")
	cc, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect :%v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f", c)
	doUnary(c)
	
}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Println("Unary called")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Bambang",
			LastName: "Subarjo",
		},
	}
	res, err := c.Greet(context.Background(),req)
	if err != nil {
		log.Fatalf("error while calling greet rpc: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}