# best-practice-go-zero

go-zero最佳实践

## How to create a service/project:

- exec `goctl api new my_zero`,then you will see a new service named my-zero.

## start with microservice

### user-rpc

- if `protoc,protoc-gen-go,protoc-gen-rpc-go` are not installed, try with `goctl env check -i -f`;
- write a *.proto eg:[user.proto](./user/rpc/user.proto)
- exec `cd ./user/rpc && goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.`
- do your business in [getUserLogic.go](./user/rpc/internal/logic/getUserLogic.go)

## Run with api file

- cd into target file && exec `goctl api go -api my_zero.api -style goZero -dir .`

## Create a Dockerfile

- cd into target file && exec `goctl docker -go my-zero.go`

## Do some compare

### migrate or autogen model?

- migrate:gorm migrate
    - auto gen db DDL through models;
    - you cannot drop a index with migrate;
    - it's convenient to code in some special columns.
- autogen model: [go-gorm gen](https://github.com/go-gorm/gen/blob/master/README.ZH_CN.md#create-record)
    - autogen model:connect to DB,and use DDL to gen model CURD && model columns;
    - you don't need care index in dbs;
    - it's convenient to sync model in code through db.
- choose the most suitable is necessary.

## TODO list

+ [ ] 集成一个比较好用的log插件
+ [ ] 使用协程池处理一些并发较高的method或者逻辑
+ [ ] rabbitmq生产者消费者优化 TODO (断线重连优化)
+ [ ] casbin权限
+ [ ] try with [dtm](https://github.com/dtm-labs/dtm)
+ [x] validate集成校验
+ [x] 华为obs集成
+ [x] rpc中etcd服务,如果使用k8s部署,那么将直接使用target注册到k8s中,由k8s的服务发现处理
+ [x] prometheus服务监控
+ [x] 集成mysql(use gorm)
+ [x] 集成mongo(zero原生支持的mgo)
+ [x] go-zero PeriodLimit => 滑动窗口实现的限流器 => 当然go-zero也有基于令牌桶实现的限流器

## 参考资料

- [go-zero实战：让微服务Go起来——2 环境搭建](https://juejin.cn/post/7036010137408143373#heading-4)
- [go-zero实战：让微服务Go起来——9 服务监控 Prometheus](https://juejin.cn/post/7044509187027501063)
- [基于Docker搭建Prometheus和Grafana](https://www.cnblogs.com/xiao987334176/p/9930517.html#autoid-3-0-0)