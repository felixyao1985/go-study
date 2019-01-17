# grpc
win 安装流程

    google.golang.org: grpc相关组件默认获取package的路径
1、安装Protobuf

    https://github.com/google/protobuf/releases

    下载 protoc-3.5.1-win32.zip

    把解压后的 protoc.exe 放入到 GOPATH\BIN 中
    
2、安装grpc

    git clone https://github.com/grpc/grpc-go
    将grpc-go更名为grpc放入到google.golang.org中

3、安装Genproto

    git clone  https://github.com/google/go-genproto
    将clone下来的文件夹更名为genproto，放到google.golang.org下  
     
4、下载text包

    为了使包的导入方式不变，需要在src目录下面构造目录结构
    mkdir -p $GOPATH/src/golang.org/x/
    cd $GOPATH/src/golang.org/x/
    git clone https://github.com/golang/text.git    
    
5、下载net包

    为了使包的导入方式不变，需要在src目录下面构造目录结构
    mkdir -p $GOPATH/src/golang.org/x/
    cd $GOPATH/src/golang.org/x/
    
    git clone https://github.com/golang/net.git net
    go install net
    执行go install之后没有提示，就说明安装好了。    
6、安装proto

    go get -u github.com/golang/protobuf/proto

7、安装protoc-gen-go

    go get -u github.com/golang/protobuf/protoc-gen-go  
    
   
    
将proto文件编译为go文件
- proto文件语法详解参阅：https://blog.csdn.net/u014308482/article/details/52958148

      // protoc --go_out=plugins=grpc:{输出目录}  {proto文件}
      protoc --go_out=plugins=grpc:./test/ ./test.proto
    
注意：原则上不要修改编译出来的*.bp.go文件的代码，因为双方接口基于同一个proto文件编译成自己的语言源码，此文件只作为接口数据处理，业务具体实现不在*.bp.go中。
    
grpc+ProtoBuf所需的一些资源 

-   1.golang.org\x\net\context

    对应的可访问链接：https://github.com/golang/net，里面包含context，dns，http2等一系列资源

-  2.golang.org/x/text/secure/bidirule

    对应的可访问链接：https://github.com/golang/text，里面包含cmd，currency，secure等一系列资源

-  3.google.golang.org/grpc

    对应的可访问链接：https://github.com/grpc/grpc-go，里面包含connectivity，grpclb，grpclog等一系列资源

-  4.google.golang.org/genproto

    对应的可访问链接：https://github.com/google/go-genproto，里面包含googleapis，protobuf等一系列资源   