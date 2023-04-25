GOBIN ?= $(shell go env GOPATH)/bin

# 运行测试用例
.PHONY:test
test:
	go test -race ./..

# 格式化项目代码
.PHONY:fmt
fmt:
	goimports -l -w .

# 检测语法 => TODO 检测lint是否安装
.PHONY:lint
lint:
	golangci-lint run

# 整理项目依赖
.PHONY:tidy
tidy:
	go mod tidy

.PHONY:base
base:
	sudo -E docker build -t zero:base -f Dockerfile .

# 构建镜像: 构建所有 make build SVC=''
.PHONY:build
build: # make build SVC='my-zero'
	sudo -E docker-compose build ${SVC}
	sudo -E docker-compose push  ${SVC}

# 重启服务:
.PHONY:start
start: # make start SVC='my-zero'
	sudo -E docker-compose stop  ${SVC}
	sudo -E docker-compose pull  ${SVC}
	sudo -E docker-compose up -d ${SVC}

# 清理docker，释放无用空间
.PHONY:clean-docker
clean-docker:
	sudo -E docker container prune -f
	sudo -E docker volume prune -f
	#sudo -E docker image prune -f

.PHONY:env
env:
	echo DOCKER_TAG="${DOCKER_TAG}" >> .env
	echo PROJECT="${PROJECT}" >> .env
	echo HUB_DOMAIN="${HUB_DOMAIN}" >> .env
	echo BRANCH_ENV="${BRANCH_ENV}" >> .env

.PHONY:api # 脚手架生成服务框架  eg:make api SVC=pay
api:
	goctl api go -api ./${SVC}/${SVC}.api -style go_zero -dir ./${SVC}

.PHONY:swagger # 根据swagger文件生成对应的client  eg:make api SVC=pay
swagger:
	swagger-codegen generate -i ./${SVC}/${SVC}.json -l go -o ./gen/${SVC}

# TODO add json