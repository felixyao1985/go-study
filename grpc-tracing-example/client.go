package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	gtrace "study/go-study/grpc-tracing-example/interceptor"
	"study/go-study/proto"
)

var (
	c_jaegerTracerServer = "106.14.125.33:6831"
	c_rpcServer          = "127.0.0.1:8028"
)

func rpcCli(dialOpts []grpc.DialOption) {
	conn, err := grpc.Dial(c_rpcServer, dialOpts...)
	if err != nil {
		fmt.Printf("grpc connect failed, err:%+v\n", err)
		return
	}
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := test.NewWaiterClient(conn)

	// 模拟请求数据
	res := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		res = os.Args[1]
	}

	// 调用gRPC接口
	tr, err := t.DoMD5(context.Background(), &test.Req{Str: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %s", tr.BackStr)
}

func main() {
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}
	tracer, _, err := gtrace.NewJaegerTracer("testCli", c_jaegerTracerServer)
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}

	if tracer != nil {
		dialOpts = append(dialOpts, gtrace.DialOption(tracer))
	}
	// do rpc-call with dialOpts
	rpcCli(dialOpts)
}
