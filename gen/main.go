package main

import (
	"context"
	"fmt"
	swagger "github.com/zjzjzjzj1874/best-pracrice-go-zero/gen/pay"
)

func main() {
	ctx := context.Background()
	cfg := swagger.NewConfiguration()
	cfg.BasePath = "http://127.0.0.1/"
	payClient := swagger.NewAPIClient(cfg)
	livesRes, resp, err := payClient.LivenessApi.Liveness(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)
	fmt.Println(livesRes)
}
