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