package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ikalkali/belajar-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)


func main() {
	fmt.Println("Calculator client!")
	cc, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cant connect : %v", err)
	}

	defer cc.Close()
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])

	firstNum := int32(a)
	secondNum := int32(b)

	c:=calculatorpb.NewSumServiceClient(cc)

	req := &calculatorpb.CalculatorRequest{
		FirstNumber: firstNum,
		SecondNumber: secondNum,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling sum rpc: %v", err)
	}
	log.Printf("Sum of %v + %v: %v", firstNum, secondNum, res.Result)
}