# 这个可以生成客户端调用的代码,但是Mac M1上架构不支持运行该命令!!!
for l in go javascript php; do
  docker run --rm -v "$(pwd):/go-work" swaggerapi/swagger-codegen-cli generate \
    -i "/go-work/swagger.json" \
    -l "$l" \
    -o "/go-work/clients/$l"
done