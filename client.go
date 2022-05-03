package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-demo/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接异常", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)
	client := pb.NewUserInfoServiceClient(conn)
	req := new(pb.UserRequest)
	req.Name = "abc"
	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("连接失败", err)
	}
	fmt.Println("连接成功", resp)
}
