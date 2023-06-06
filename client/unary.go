package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MilkiasTonji/grpc-test/proto"
)

func CallSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Failed to fetch Message %v", err)
	}

	log.Printf("%s", res.Message)
}
