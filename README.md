# 悟空IM GO SDK

使用WuKongIM的gosdk，在原有的WuKongIMGoSDK基础上进行改进，不喜欢使用大小写混用命名，所以改名。

参照examples/example.go的调用，使用如下编译

```bash
go build -o bin_exam examples/example.go
```

引入包

```go
go get github/yytt5301/wkim_gosdk
```

使用
```go
import  "github/yytt5301/wkim_gosdk/pkg/wksdk"
```

初始化

```go
cli := wksdk.NewClient("tcp://127.0.0.1:5100",wksdk.WithUID("xxxx"),wksdk.WithToken("xxxxx"),wksdk.WithReconnect(true))
```

连接

```go

cli.Connect()

// 监听连接状态
cli.OnConnect(func(status Status,reasonCode int) {
    //TODO
})

```

发送消息

```go

result,_ := cli.SendMessage(content,channel)

```

接受消息

```go

cli.OnMessage(func(m *Message) {
    //TODO
})


```