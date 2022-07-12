# go-zero的镜像是多阶段构建的,这个是构建镜像,下面还有个FROM是运行时镜像;
# 下面使用 `COPY --from=builder /app/etc /app/etc` 从上一阶段的镜像中复制文件，我们也可以复制任意镜像中的文件。
# 请参考 https://yeasy.gitbook.io/docker_practice/image/multistage-builds
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./my_zero/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/my-zero my_zero/my-zero.go

# 不允许进入容器就用scratch
#FROM scratch
# 需要进入容器用这个镜像alpine
FROM alpine


COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/my-zero /app/my-zero
COPY --from=builder /app/etc /app/etc

CMD ["./my-zero", "-f", "etc/my-zero.yaml"]