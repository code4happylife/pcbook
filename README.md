# pcbook
grpc codes from:https://www.bilibili.com/video/BV1Rg411T7yF?p=8&amp;share_source=copy_web

# protobuf

编写 .proto 文件来描述要传输的对象，protoc编译器会把它编译成不同语言对该类对象进行序列化和反序列化的API

# gRPC

gRPC由google开发，是一款语言中立、平台中立、开源的远程过程调用系统,基于HTTP2协议标准设计开发

gRPC可以实现微服务，将大的项目拆分为多个小且独立的业务模块，也就是服务，各服务间使用高效的protobuf协议进行RPC调用，gRPC默认使用protocol buffers

# 完整的RPC流程


客户端（gRPC Sub）调用 A 方法，发起 RPC 调用

对请求信息使用 Protobuf 进行对象序列化压缩（IDL）

服务端（gRPC Server）接收到请求后，解码请求体，进行业务逻辑处理并返回

对响应结果使用 Protobuf 进行对象序列化压缩（IDL）

客户端接受到服务端响应，解码请求体。回调被调用的 A 方法，唤醒正在等待响应（阻塞）的客户端调用并返回响应结果

# 项目结构

cmd 目录下分别存在client 端和 server 端的代码, 在超时场景下的输出：

server端输出：
```
pcbook % make server
go run cmd/server/main.go -port 8080
2022/01/16 16:39:03 start server on port:8080
2022/01/16 16:39:09 receive a create-laptop request with id:98981717-beb4-4d8c-938c-cc19b8674b72
2022/01/16 16:39:15 deadline is exceeded

```

client端输出：
```
pcbook % make client
go run cmd/client/main.go -address 0.0.0.0:8080
2022/01/16 20:57:18 Dial server 0.0.0.0:8080
2022/01/16 20:57:23 can not create laptop:rpc error: code = DeadlineExceeded desc = context deadline exceeded
exit status 1
make: *** [client] Error 1

```
