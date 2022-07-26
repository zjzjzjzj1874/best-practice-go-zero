# 运行测试用例
test:
	go test -race ./..

# 格式化项目代码
fmt:
	goimports -l -w .

# 检测语法 => TODO 检测lint是否安装
lint:
	golangci-lint run

# 整理项目依赖
tidy:
	go mod tidy

# 构建所有服务 - 基于docker-compose
all:
	docker-compose build
	docker-compose push

# 构建my-zero:
my-zero:
	docker-compose build my-zero
	docker-compose push my-zero

# 重启my-zero:
my-zero-start:
	docker-compose stop my-zero
	#docker-compose pull my-zero
	docker-compose up -d my-zero