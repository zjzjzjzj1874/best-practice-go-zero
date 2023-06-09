- 重新生成客户端:当api文件有改动时
  `goctl api go -api task.api -style go_zero -dir .`
- 重新构建镜像
  `docker build -t xx .`
- 手动调用定时任务
  `curl -X POST http://localhost:30008/task/v0/manual -d '{"name":"测试任务"}' --header "Content-Type: application/json"`