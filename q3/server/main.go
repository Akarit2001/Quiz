package main

import (
	"fmt"
	"log"
	"net"

	"github.com/akarit2001/quiz/q3/server/services"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

// grpc
func main() {
	// test output
	// fmt.Println(string(services.GetBeefs()))

	// httpServer()
	grpcServer()

}

func grpcServer() {
	s := grpc.NewServer()

	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatal(err)
	}

	services.RegisterBeefServer(s, services.NewBeefServer())

	fmt.Println("grpc server listening on port 8000")
	if err = s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func httpServer() {
	app := fiber.New()

	app.Get("/beef/summary", func(c *fiber.Ctx) error {
		return c.Send(services.GetBeefs())
	})
	fmt.Println("http server listening on port 8001")
	app.Listen(":8081")
}
