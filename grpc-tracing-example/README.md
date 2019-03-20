###基于jaeger和tracing实现 grpc调用链跟踪模块

#####Jaeger

[参考 Jaeger部署]("https://blog.csdn.net/moxiaomomo/article/details/80723758")

Jaeger是Uber推出的一款调用链追踪系统，类似于Zipkin和Dapper，为微服务调用追踪而生。 其主要用于多个服务调用过程追踪分析，图形化服务调用轨迹，便于快速准确定位问题。
Jaeger组成

- 前端界面展示UI
- 数据存储Cassandra
- 数据查询Query
- 数据收集处理Collector
- 客户端代理Agent
- 客户端库jaeger-client-*



```yaml
    
  仅为测试，该方式数据会存入内存，不能用于生产
  
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.9
```

#####grpc-jaeger源依赖
    
    google.golang.org/grpc
    github.com/opentracing/opentracing-go
    github.com/uber/jaeger-client-go
    github.com/pkg/errors
