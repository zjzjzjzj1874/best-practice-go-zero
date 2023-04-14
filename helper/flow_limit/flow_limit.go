package flow_limit

type Conf struct {
	PeriodSec int `json:",default=3600"` // 限流周期(s) 默认3600
	Quota     int `json:",default=100"`  // 限流频次 默认100次
}
