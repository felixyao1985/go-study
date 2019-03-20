package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"study/go-study/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	t := test.NewWaiterClient(conn)
	res := "test123"

	// 调用gRPC接口
	tr, err := t.DoMD5(context.Background(), &test.Req{Str: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %s", tr.BackStr)
}
