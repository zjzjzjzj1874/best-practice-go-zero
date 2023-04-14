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
COPY ./pay/etc /app/etc
# json文件写入
COPY ./pay/pay.json /app/
RUN go build -ldflags="-s -w" -o /app/pay pay/pay.go

FROM alpine

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/pay /app/pay
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/pay.json /app/

CMD ["./pay", "-f", "etc/pay-api.yaml"]