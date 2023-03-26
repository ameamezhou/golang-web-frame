# 独立路由 和 Context
### 一、 封装 response 和 request --> Context
操作 `response` 和 `request` 太麻烦，并且需要构造和设置的东西太多
> 如果我们要构建一个完整的响应，肯定需要考虑消息头`Header`和消息体`Body`.
> 而Header包含了状态码`StatusCode`和消息类型`ContentType`等几乎每次请求都需要设置的消息.
> 所以如果我们不进行有效的封装，那么框架的用户将需要写大量重复，发杂的代码，并且还很容易出错
> 我们需要针对使用场景，能够高效地构造出HTTP响应是一个好框架必须考虑的点。

这里的代码参考 base3/compare/compare.go

