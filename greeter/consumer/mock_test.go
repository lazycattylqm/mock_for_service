package main

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	greetMock "liqiming.com.greeter/consumer/mock_greeter"
	pb "liqiming.com.greeter/proto"
	"testing"
)

func TestMock(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	service := greetMock.NewMockGreeterService(ctl)
	service.EXPECT().SayHi(gomock.Any(), &pb.FromReq{ // 这里可用任何入参表明mock结果
		Name: "liqiming",
	}).Return(&pb.ResMsg{ // return 我们要的结果
		Name: "mock",
	}, nil) // 这里可
	hi, err := service.SayHi(
		context.Background(), &pb.FromReq{
			Name: "liqiming",
		},
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hi.Name)

	service.EXPECT().SayHi(gomock.Any(), gomock.Any()).Return(&pb.ResMsg{
		Name: "23444",
	}, nil)
	hi, err = service.SayHi(
		context.Background(), &pb.FromReq{
			Name: "1243",
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hi.Name)

}
