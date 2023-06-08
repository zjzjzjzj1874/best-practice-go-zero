PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))
VERSION = v$(shell cat .version)
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

GOBUILD=CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

GOINSTALL=CGO_ENABLED=0 go install -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

GOBIN ?= $(shell go env GOPATH)/bin

.PHONY:echo
echo:
	@echo "PKG:${PKG}"
	@echo "VERSION:${VERSION}"
	@echo "COMMIT_SHA:${COMMIT_SHA}"
	@echo "GOOS:${GOOS}"
	@echo "GOARCH:${GOARCH}"
	@echo "GOBUILD:${GOBUILD}"
	@echo "GOINSTALL:${GOINSTALL}"
	@echo "GOBIN:${GOBIN}"

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
api: $(GOBIN)/goctl
	goctl api go -api ./${SVC}/${SVC}.api -style go_zero -dir ./${SVC}

.PHONY:swagger # 根据swagger文件生成对应的client  eg:make api SVC=pay
swagger:  # 注意:command -v swagger-codegen 用于查找swagger-codegen的绝对路径;
	@if ! [ -x "$(command -v swagger-codegen)" ]; then \
		echo 'Installing swagger-codegen...'; \
		brew install swagger-codegen; \
	fi
	swagger-codegen generate -i ./${SVC}/${SVC}.json -l go -o ./gen/${SVC}

.PHONY:json # eg:make json SVC=pay
json: $(GOBIN)/goctl # 这里表示需要检查是否已经安装了goctl工具,没有的话要先安装
	goctl api plugin -plugin goctl-swagger="swagger -filename ./${SVC}/${SVC}.json" -api ./${SVC}/${SVC}.api -dir .

.PHONY:dockerfile # eg:make dockerfile SVC=pay
dockerfile: $(GOBIN)/goctl
	cd ./${SVC} && goctl docker -go ${SVC}.go && cd -

.PHONY:kube # eg:make kube SVC=pay
kube: $(GOBIN)/goctl
	goctl kube deploy -name backend-${SVC} -namespace default  -image ${SVC} -o ./deploy/my-zero/${SVC}-k8s.yaml -port 80 --serviceAccount find-endpoints

$(GOBIN)/goctl:
	go install github.com/zeromicro/go-zero/tools/goctl@latest