# easycs
一个令http客户端变的极简的库，不必再创建request，在判断一堆错误码，也不用再defer response.Body.Close()，一行实现http请求

## 安装
```go
import "github.com/bournex/easycs"
```

## 用法

### 请求 http://localhost:8000
```go
var ec easycs.EasyC
response, err := ec.Do()
if err != nil{
    // do something
}
```

### 请求 http://172.16.32.58:8888/
```go
ec.WithHost("172.16.32.58").WithPort("8888").Do()
```

### 同步请求 POST http://www.example.com/dosomething
```go
ec.WithScheme("POST").WithUrl("http://www.example.com/dosomething").WithForm("user", "allen").Do()
```

### 带回调的同步请求 POST http://www.example.com/dosomething
```go
ec.WithScheme("POST").WithUrl("http://www.example.com/dosomething").WithForm("user", "allen").Do(
    func (r *easycs.Response, err){
        fmt.Printf("status: %d, body %s\n", r.Status, string(r.Body))
    }
)
```

### 异步请求 POST http://www.example.com/dosomething
```go
ec.WithScheme("POST").WithUrl("http://www.example.com/dosomething").WithForm("user", "allen").Done(
    func (r *http.Response, err error)){
        fmt.Printf("status: %d, body %s\n", r.Status, string(r.Body))
    }
)
```

### 带回调的异步请求 POST http://www.example.com/dosomething
```go
ec.WithScheme("POST").WithUrl("http://www.example.com/dosomething").WithForm("user", "allen").DoneWithStatus(
    func (r *easycs.Response, err){
        fmt.Printf("status: %d, body %s\n", r.Status, string(r.Body))
    }
)
```