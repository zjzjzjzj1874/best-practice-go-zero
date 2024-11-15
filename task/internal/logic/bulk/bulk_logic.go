package bulk

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"sync"

	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/svc"
	"github.com/zjzjzjzj1874/best-pracrice-go-zero/task/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BulkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBulkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BulkLogic {
	return &BulkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Bulk 参考其他的实现
func (l *BulkLogic) Bulk(req *types.BulkReq) (resp *types.BulkResp, err error) {
	// go-zero中 BulkExecutor的使用
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// 创建 BulkExecutor
	//executor := executors.NewBulkExecutor()
	//testCh := make(chan *model.TTest, 10)
	threading.GoSafe(func() {
		defer wg.Done()

	})

	return
}
