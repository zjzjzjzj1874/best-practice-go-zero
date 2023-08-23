# best-practice-go-zero

go-zero最佳实践

## 项目结构

```shell
.
├── Dockerfile -- docker构建文件
├── Jenkinsfile -- jenkins流水线脚本
├── LICENSE -- MIT许可证
├── Makefile 
├── README.md -- 文档
├── __test__ -- 测试单元
├── constants -- 常量
│   ├── errors -- 错误定义
│   └── types -- 枚举类型定义
├── deployment -- 部署文件夹
├── docker-compose.yml
├── example -- 示例代码
├── go.mod -- golang依赖包
├── helper -- 工具类
│   ├── area.go -- 中国省市区解析器
│   ├── area_test.go -- 单元测试
│   ├── breaker -- 限流器
│   ├── cmd.go -- 命令行执行器
│   ├── contextx -- 上下文
│   ├── converter.go -- 类型转换器
│   ├── division.go -- sqlite的区域
│   ├── division_test.go -- 单测
│   ├── divisions.sqlite -- 地区文件sqlite
│   ├── email -- 邮件工具类
│   ├── ffmpeg.go -- FFmpeg执行
│   ├── ffmpeg_test.go -- 单测
│   ├── logrusx.go -- 日志中间件
│   ├── obs -- 华为对象存储
│   ├── pprof.go -- pprof工具
│   ├── pprof_test.go -- 单测
│   ├── rabbitmq -- 消息队列(rabbitmq)
│   ├── redis.go -- redis中间件
│   ├── redis_model.go -- 模型定义
│   ├── redis_test.go -- 单测
│   ├── sqlx -- 数据库工具类
│   ├── swagger.go -- swagger工具
│   ├── tracex -- 链路追踪工具
│   ├── validate.go -- validate工具
│   └── validate_test.go -- 单测
├── middlewares -- 中间件
│   └── log_trace.go -- 日志中间件
├── model -- 数据库模型
│   ├── mongo -- MongoDB模型
│   └── mysql -- MYSQL模型
├── my_zero -- zero模块
│   ├── Dockerfile -- 构建文件
│   ├── README.md -- 文档
│   ├── api -- api文件夹
│   ├── etc -- 配置文件
│   ├── internal -- 内部逻辑
│   ├── my-zero.go -- 入口函数
│   └── my_zero.api -- 入口api文件
├── order -- 订单模块
│   └── api -- api文件夹
├── scripts -- 脚本
│   ├── k8s_deploy.sh -- 部署脚本
├── static -- 静态文件
├── task -- 任务模块
│   ├── Dockerfile -- 构建文件
│   ├── README.md -- 文档
│   ├── api -- api文件夹
│   ├── etc -- 配置文件
│   ├── internal -- 内部逻辑
│   ├── swagger.json -- swagger文件
│   ├── task-k8s.yaml -- 部署文件
│   ├── task.api -- 入口api
│   ├── task.go -- 入口函数
│   └── task.sh -- 脚本
├── template -- goctl模板
│   └── mongo -- mongo模板
├── user -- 用户模块
│   └── rpc -- rpc模块
```

## go-zero usage
### How to create a service/project:

- exec `goctl api new my_zero`,then you will see a new service named my-zero.

### start with microservice

#### user-rpc

- if `protoc,protoc-gen-go,protoc-gen-rpc-go` are not installed, try with `goctl env check -i -f`;

- write a \*.proto eg:[user.proto](./user/rpc/user.proto)

- exec `cd ./user/rpc && goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.`

- do your business in [getUserLogic.go](./user/rpc/internal/logic/getUserLogic.go)

### Run with api file

- cd into target file && exec `goctl api go -api my_zero.api -style goZero -dir .`

### Create a Dockerfile

- cd into target file && exec `goctl docker -go my-zero.go`

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

### `goctl`实用小技巧

- 根据api文件生成接口:`goctl api go -api task.api -style go_zero -dir .`,驼峰命名的话`go_zero`替换成`goZero`

- 根据api文件生成swagger:`goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api task.api`

- Dockerfile生成:`goctl docker -go task.go`

- K8S部署yaml生成:`goctl kube deploy -name task -namespace my-ns  -image task -o task-k8s.yaml -port 80 --serviceAccount find-endpoints`

- 生成`Mongo Model`:`goctl model mongo -type Task -c -style go_zero -d .`

## TODO list

+ [x] 集成一个比较好用的log插件:目前使用logrus

+ [x] Go-Zero自带链路追踪:Jaeger:\[配置位置\](./task/etc/task.yaml),仅需要在配置中添加\`Telemetry\`信息.[配置位置](./task/etc/task.yaml),仅需要在配置中添加`Telemetry`信息.

+ [x] Opentelemetry+Jaeger链路追踪:\[代码位置\](./example/otel)[代码位置](./example/otel)

+ [ ] 链路追踪+日志追踪:链路追踪有traceId,如果使用中间件获取每个request中的traceId,然后使用logrus的hook,日志打印过程中有traceId则添加,链路追踪会更加完善.

+ [ ] 使用协程池处理一些并发较高的method或者逻辑

+ [ ] rabbitmq生产者消费者优化 TODO (断线重连优化)

+ [ ] casbin权限

+ [ ] try with \[dtm\](https://github.com/dtm-labs/dtm)[dtm](https://github.com/dtm-labs/dtm)

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

+ [x] 使用swagger文件生成golang客户端代码 => \`swagger-codegen generate -i swagger.json -l go -o ./gen\`,具体使用见下方\`根据swagger文件生成golang客户端\``swagger-codegen generate -i swagger.json -l go -o ./gen`,具体使用见下方`根据swagger文件生成golang客户端`

+ [x] json优化:1.json.NewEncoder代替json.marshal;2.使用json.Encoder的底层缓冲区,减少内存分配和垃圾回收开销
+ [ ] 有http请求的路由中新增log中间件,用于请求的path和消耗的时间;
+ [ ] TODO 写一个类似[gptx](https://github.com/zjzjzjzj1874/chatgpt)的wechat发消息的命令行工具;加油!!!

## project deploy
### 部署步骤

以task服务为例

- 生成task的`Dockerfile`:`goctl docker -go task.go`

- `docker-compose`中添加task服务:然后:`docker-compose build task`

### api脚手架生成代码

- pay

```shell
make api SVC=pay
```

- task

```shell
make api SVC=task
```

### swagger json文件生成

- pay

```shell
make json SVC=pay
```

- task

```shell
make json SVC=task
```

### 根据swagger文件生成golang客户端

- pay

```shell
make swagger SVC=pay
```

### Dockerfile生成

- pay

```shell
make dockerfile SVC=pay
```

## 其他依赖

### 微信支付
- [go-client](https://github.com/wechatpay-apiv3/wechatpay-go)
- [postman脚本](https://github.com/wechatpay-apiv3/wechatpay-postman-script)

## 参考资料

- [go-zero实战：让微服务Go起来——2 环境搭建](https://juejin.cn/post/7036010137408143373#heading-4)

- [go-zero实战：让微服务Go起来——9 服务监控 Prometheus](https://juejin.cn/post/7044509187027501063)

- [基于Docker搭建Prometheus和Grafana](https://www.cnblogs.com/xiao987334176/p/9930517.html#autoid-3-0-0)


### 其他说明

go-zero在处理multipart-formdata的结构时,如果path中有参数,formdata中也有参数,请参考[import](./task/internal/handler/imp/import_handler.go)代码中的处理方法.