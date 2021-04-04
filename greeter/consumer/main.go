package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3"
	greeter "liqiming.com.greeter/proto"
)

func main() {
	service := micro.NewService()
	service.Init()
	greeterServiceClient := greeter.NewGreeterService("greeter.provider", service.Client())

	hi, err := greeterServiceClient.SayHi(
		context.TODO(), &greeter.FromReq{
			Name: "liqiming",
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(hi.Name)

}
