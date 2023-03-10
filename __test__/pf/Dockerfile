FROM golang as buiblder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /
ADD go.mod .

COPY . .
RUN go build -o /test ./pf.go

FROM alpine

WORKDIR /
COPY --from=buiblder test .
CMD ["/test"]