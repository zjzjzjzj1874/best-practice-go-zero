FROM golang:alpine AS builder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY ./task/etc /app/etc
COPY ./task/swagger.json /app/
RUN go build -ldflags="-s -w" -o /app/task task/task.go

FROM alpine

WORKDIR /app
COPY --from=builder /app/task /app/task
COPY --from=builder /app/etc /app/etc
COPY --from=builder /app/swagger.json /app/

CMD ["./task", "-f", "etc/task.yaml"]