package breaker

import (
	"errors"
	"sync"
	"time"
)

type CircuitBreaker struct {
	sync.Mutex

	// 熔断器状态
	state State

	// 熔断器参数
	maxRequests      uint32        // 最大请求量
	intervalDuration time.Duration // 统计时间段
	timeoutDuration  time.Duration // 超时时间

	// 熔断器计数器
	windowStart time.Time  // 窗口开始时间
	requests    []*Request // 请求记录

	// 熔断器指标
	consecutiveFailures uint32 // 连续失败次数
	totalFailures       uint32 // 总失败次数
	totalSuccesses      uint32 // 总成功次数

	// 熔断器回调函数
	onStateChange func(State)
}
type Request struct {
	StartTime time.Time // 请求开始时间
}
type State uint8

const (
	StateClosed State = iota
	StateHalfOpen
	StateOpen
)

func NewCircuitBreaker(maxRequests uint32, intervalDuration, timeoutDuration time.Duration) *CircuitBreaker {
	cb := &CircuitBreaker{
		maxRequests:      maxRequests,
		intervalDuration: intervalDuration,
		timeoutDuration:  timeoutDuration,
		state:            StateClosed,
		onStateChange:    func(State) {},
	}
	go cb.monitor()
	return cb
}

func (cb *CircuitBreaker) monitor() {
	ticker := time.NewTicker(cb.intervalDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			cb.Lock()
			if cb.state == StateHalfOpen {
				cb.state = StateClosed
				cb.consecutiveFailures = 0
				cb.totalFailures = 0
			}
			cb.requests = nil
			cb.windowStart = time.Now()
			cb.Unlock()
		}
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.Lock()
	defer cb.Unlock()

	if cb.state == StateOpen {
		if time.Now().Sub(cb.windowStart) > cb.timeoutDuration {
			cb.state = StateHalfOpen
			cb.onStateChange(cb.state)
		} else {
			return errors.New("circuit breaker is open")
		}
	}

	if cb.state == StateHalfOpen {
		return cb.try(fn)
	}

	if len(cb.requests) < int(cb.maxRequests) {
		return cb.try(fn)
	}

	if cb.failureRatio() >= 0.5 {
		cb.state = StateOpen
		cb.windowStart = time.Now()
		cb.onStateChange(cb.state)
		return errors.New("circuit breaker is open")
	}

	return cb.try(fn)
}

func (cb *CircuitBreaker) try(fn func() error) error {
	request := &Request{StartTime: time.Now()}
	cb.requests = append(cb.requests, request)
	cb.Unlock()
	err := fn()

	cb.Lock()
	requests := cb.requests
	cb.totalSuccesses++

	// 删除过期的请求记录
	expiryTime := time.Now().Add(-cb.intervalDuration)
	var i int
	for i = range requests {
		if requests[i].StartTime.Before(expiryTime) {
			break
		}
	}
	if i > 0 {
		cb.requests = requests[i:]
	}

	// 统计熔断器指标
	if err != nil {
		cb.consecutiveFailures++
		cb.totalFailures++
		if cb.consecutiveFailures >= cb.maxRequests {
			cb.state = StateOpen
			cb.windowStart = time.Now()
			cb.onStateChange(cb.state)
			return errors.New("circuit breaker is open")
		}
	} else {
		cb.consecutiveFailures = 0
	}

	if cb.state == StateHalfOpen {
		cb.state = StateClosed
		cb.onStateChange(cb.state)
	}

	return err
}
func (cb *CircuitBreaker) failureRatio() float64 {
	if cb.totalFailures == 0 {
		return 0.0
	}
	return float64(cb.totalFailures) / float64(cb.totalSuccesses+cb.totalFailures)
}

func (cb *CircuitBreaker) State() State {
	cb.Lock()
	defer cb.Unlock()
	return cb.state
}

func (cb *CircuitBreaker) SetState(s State) {
	cb.Lock()
	cb.state = s
	cb.Unlock()
}

func (cb *CircuitBreaker) OnStateChange(fn func(State)) {
	cb.Lock()
	cb.onStateChange = fn
	cb.Unlock()
}

func (cb *CircuitBreaker) FailureCount() uint32 {
	cb.Lock()
	defer cb.Unlock()
	return cb.totalFailures
}

func (cb *CircuitBreaker) SuccessCount() uint32 {
	cb.Lock()
	defer cb.Unlock()
	return cb.totalSuccesses
}

func (cb *CircuitBreaker) ConsecutiveFailureCount() uint32 {
	cb.Lock()
	defer cb.Unlock()
	return cb.consecutiveFailures
}

func (cb *CircuitBreaker) Reset() {
	cb.Lock()
	cb.requests = nil
	cb.windowStart = time.Now()
	cb.consecutiveFailures = 0
	cb.totalFailures = 0
	cb.totalSuccesses = 0
	cb.state = StateClosed
	cb.Unlock()
}

func (cb *CircuitBreaker) IsOpen() bool {
	return cb.State() == StateOpen
}

func (cb *CircuitBreaker) IsClosed() bool {
	return cb.State() == StateClosed
}

func (cb *CircuitBreaker) IsHalfOpen() bool {
	return cb.State() == StateHalfOpen
}
