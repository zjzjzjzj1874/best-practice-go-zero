# best-practice-go-zero

go-zero最佳实践

## How to create a service/project:

- exec `goctl api new my_zero`,then you will see a new service named my-zero.

## Run with api file

- cd into target file && exec `goctl api go -api my_zero.api -style goZero -dir .`

## Create a Dockerfile

- cd into target file && exec `goctl docker -go my-zero.go`

## TODO list

- validate集成校验
- 集成一个比较好用的log插件
- 使用协程池处理一些并发较高的method或者逻辑
- rabbitmq生产者消费者优化 TODO (断线重连优化)