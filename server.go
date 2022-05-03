package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
	"net"
)

type UserInfoService struct {
}

var u = UserInfoService{}

func (service *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	if name == "abc" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   19,
			Hobby: []string{"hello", "world"},
		}
	}
	err = nil
	return
}

func main() {
	address := "127.0.0.1:8080"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("监听异常")
		return
	}
	fmt.Println("监听成功")
	s := grpc.NewServer()
	pb.RegisterUserInfoServiceServer(s, &u)
	_ = s.Serve(lis)
}
