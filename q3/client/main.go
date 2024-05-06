package main

import (
	"context"
	"fmt"
	"log"

	"github.com/akarit2001/quiz/q3/client/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	cc, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()
	client := services.NewBeefClient(cc)
	res, err := client.Summary(context.Background(), &services.EmptyRequest{})
	if err != nil {
		log.Panicln(res)
		return
	}
	fmt.Println(string(res.Data))

}
