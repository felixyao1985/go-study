package main

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	gtrace "study/go-study/grpc-tracing-example/interceptor"
	"study/go-study/proto"
)

type server struct{}

var (
	jaegerTracerServer = "106.14.125.33:6831"
	rpcServer          = "127.0.0.1:8028"
)

func (s *server) DoMD5(ctx context.Context, in *test.Req) (*test.Res, error) {
	fmt.Println("MD5方法请求JSON:" + in.Str)
	return &test.Res{BackStr: "MD5 :" + fmt.Sprintf("%x", md5.Sum([]byte(in.Str)))}, nil
}

func main() {
	var servOpts []grpc.ServerOption
	tracer, _, err := gtrace.NewJaegerTracer("testSrv", jaegerTracerServer)
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	if tracer != nil {
		servOpts = append(servOpts, gtrace.ServerOption(tracer))
	}

	s := grpc.NewServer(servOpts...)

	test.RegisterWaiterServer(s, &server{})
	reflection.Register(s)

	ln, err := net.Listen("tcp", rpcServer)
	if err != nil {
		os.Exit(-1)
	}
	err = s.Serve(ln)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
