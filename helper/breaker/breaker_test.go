package breaker

import (
	"errors"
	"testing"
	"time"
)

func TestCircuitBreaker(t *testing.T) {
	maxRequests := uint32(5)
	intervalDuration := time.Second
	timeoutDuration := time.Second * 10

	cb := NewCircuitBreaker(maxRequests, intervalDuration, timeoutDuration)

	// 模拟调用远程服务
	var counter uint32
	fn := func() error {
		counter++
		if counter <= 3 {
			return errors.New("remote service error")
		}
		return nil
	}

	// 连续调用远程服务 5 次，触发熔断器
	for i := 0; i < 5; i++ {
		err := cb.Execute(fn)
		if err != nil {
			t.Logf("execute failed: %v", err)
		}
	}

	// 等待 1 秒钟，让熔断器进入 HalfOpen 状态
	time.Sleep(intervalDuration)

	// 再次调用远程服务，预期返回 nil
	err := cb.Execute(fn)
	if err != nil {
		t.Errorf("execute failed: %v", err)
	}

	// 再次调用远程服务，预期返回错误
	err = cb.Execute(fn)
	if err == nil {
		t.Errorf("expect error, but got nil")
	}
}
