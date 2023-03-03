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

+ [x] 集成一个比较好用的log插件:目前使用logrus
+ [x] Opentelemetry+Jaeger链路追踪:代码位置`./example/otel`
+ [ ] 链路追踪+日志追踪:链路追踪有traceId,如果使用中间件获取每个request中的traceId,然后使用logrus的hook,日志打印过程中有traceId则添加,链路追踪会更加完善.
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
+ [x] go-zero自动生成swagger文件:
  - 切换到task项目并执行:`goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api task.api`
  - 如果没有安装`goctl-swagger`,请先安装`goctl-swagger`(用于生成Swagger文档的工具):`go install github.com/zeromicro/goctl-swagger@latest`
  - `export PATH=$PATH:$(go env GOPATH)/bin` --> 将所有`gopath/bin`下面的工具添加到全局变量中;
  - `source ./zshrc` --> 我用的oh my zsh,然后重新source即可
  - 服务器或者本地安装swagger-ui,然后查看网页.`docker run -it -d --name swagger-ui -p 8080:8080 swaggerapi/swagger-ui`
  - 查看`swagger.json`文件: `curl http://localhost:8888/task/swagger` ,先把`swagger.json`文件复制到镜像中,然后使用`go:embed`把二进制文件读取出来,也可以使用`ioutils.Readfile`;最后返回二进制文件流.
+ [ ] kafka客户端实现
+ [ ] elasticsearch客户端实现

## `goctl`实用小技巧
- 根据api文件生成接口:`goctl api go -api task.api -style go_zero -dir .`,驼峰命名的话`go_zero`替换成`goZero`
- 根据api文件生成swagger:`goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api task.api`
- Dockerfile生成:`goctl docker -go task.go`
- K8S部署yaml生成:`goctl kube deploy -name task -namespace my-ns  -image task -o task-k8s.yaml -port 80 --serviceAccount find-endpoints`
- 生成`Mongo Model`:`goctl model mongo -type Task -c -style go_zero -d .`

## 部署步骤
以task服务为例
- 生成task的`Dockerfile`:`goctl docker -go task.go`
- `docker-compose`中添加task服务:然后:`docker-compose build task`

## 参考资料

- [go-zero实战：让微服务Go起来——2 环境搭建](https://juejin.cn/post/7036010137408143373#heading-4)
- [go-zero实战：让微服务Go起来——9 服务监控 Prometheus](https://juejin.cn/post/7044509187027501063)
- [基于Docker搭建Prometheus和Grafana](https://www.cnblogs.com/xiao987334176/p/9930517.html#autoid-3-0-0)