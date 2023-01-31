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

base:
	sudo -E docker build -t zero:base -f Dockerfile .

# 构建镜像: 构建所有 make build SVC=''
build: # make build SVC='my-zero'
	sudo -E docker-compose build ${SVC}
	sudo -E docker-compose push  ${SVC}

# 重启服务:
start: # make start SVC='my-zero'
	sudo -E docker-compose stop  ${SVC}
	sudo -E docker-compose pull  ${SVC}
	sudo -E docker-compose up -d ${SVC}

# 清理docker，释放无用空间
clean-docker:
	sudo -E docker container prune -f
	sudo -E docker volume prune -f
	#sudo -E docker image prune -f

env:
	echo DOCKER_TAG="${DOCKER_TAG}" >> .env
	echo PROJECT="${PROJECT}" >> .env
	echo HUB_DOMAIN="${HUB_DOMAIN}" >> .env
	echo BRANCH_ENV="${BRANCH_ENV}" >> .env
