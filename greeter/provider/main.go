package main

import (
	"github.com/asim/go-micro/v3"
	pb "liqiming.com.greeter/proto"
	"liqiming.com.greeter/provider/handler"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter.provider"),
	)

	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
