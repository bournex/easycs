# easycs
一个令http客户端变的极简的库，不必再创建request，在判断一堆错误码，也不用再defer response.Body.Close()，一行实现http请求

## 安装
```go
go get -v "github.com/bournex/easycs"
```

## 接口

请求构造阶段的方法  
|方法|示例|说明|
|---|---|:---|
|WithMethod|GET|设置http方法|
|WithScheme|https|设置scheme，支持http、https|
|WithHost|www.example.com:8080|设置host，包含端口|
|WithPath|/status|设置请求路径|
|WithUrl|https://www.example.com:8080/status|设置完整路径，将忽略Scheme、Host、Path配置|
|WithClient|——|允许使用自定义的http.Client|
|WithQuery/WithQuerys|——|设置query参数|
|WithForm/WithForms|——|设置form参数|
|WithHeader/WithHeaders|——|设置请求头|
|WithContext|——|设置需要的上下文|

请求方法  
|方法|说明|
|---|:---|
|Do|发起同步http/https请求，返回http.Response和错误码|
|DoWithStatus|发起同步http/https请求，在回调中返回easycs.Response和错误码|
|Done|发起异步http/https请求，在回调中返回http.Response和错误码|
|DoneWithStatus|发起异步http/https请求，在回调中返回easycs.Response和错误码|

EasyC非线程安全，不要多线程使用同一个EasyC对象


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
ec.WithMethod("POST").WithUrl("http://www.example.com/dosomething").WithForm("user", "allen").Do()
```

### 带回调的同步请求 POST http://www.example.com/dosomething
```go
ec.WithMethod("POST").WithUrl("http://www.example.com/dosomething").WithForm("user", "allen").Do(
    func (r *easycs.Response, err){
        fmt.Printf("status: %d, body %s\n", r.Status, string(r.Body))
    }
)
```

### 异步请求 POST http://www.example.com/dosomething
```go
ec.WithMethod("POST").WithUrl("http://www.example.com/dosomething").WithForm("user", "allen").Done(
    func (r *http.Response, err error)){
        fmt.Printf("status: %d, body %s\n", r.Status, string(r.Body))
    }
)
```

### 带回调的异步请求 POST http://www.example.com/dosomething
```go
ec.WithMethod("POST").WithUrl("http://www.example.com/dosomething").WithForm("user", "allen").DoneWithStatus(
    func (r *easycs.Response, err){
        fmt.Printf("status: %d, body %s\n", r.Status, string(r.Body))
    }
)
```