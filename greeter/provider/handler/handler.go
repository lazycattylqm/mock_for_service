package handler

import (
	"context"
	"fmt"
	pb "liqiming.com.greeter/proto"
)

type Greeter struct{}

func (h *Greeter) SayHi(ctx context.Context, in *pb.FromReq, out *pb.ResMsg) error {
	out.Name = fmt.Sprintf("%v %v", "hello", in.Name)
	return nil
}
