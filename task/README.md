- 重新生成客户端:当api文件有改动时
  `goctl api go -api task.api -style go_zero -dir .`
- 重新构建镜像
  `docker build -t xx .`