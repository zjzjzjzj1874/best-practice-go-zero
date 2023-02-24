package helper

const (
	// TaskPersonalPoolPrefix 任务池前缀 :TaskPersonal+环境+用户ID => TaskPersonal.dev.123456
	TaskPersonalPoolPrefix = "TaskPersonal.%s.%s"
	// TimeoutTaskQueuePrefix 队列超时任务前缀 :TaskQueue+环境 => TaskQueue.dev
	TimeoutTaskQueuePrefix = "TimeoutTaskQueue.%s"
	// CronTaskLockerPrefix 自动任务分布式事务锁 :CronTaskLocker.环境.任务名称 => TaskQueue.dev.测试任务
	CronTaskLockerPrefix = "CronTaskLocker.%s.%s"
)

// CacheTaskQueueMetaData 清理队列缓存用户任务ID列表 => 说明: TaskQueuePrefix 表示任务的整个
type CacheTaskQueueMetaData struct {
	TaskID   string `json:"task_id"`   // 任务ID
	ExpireAt int64  `json:"expire_at"` // 过期时间戳
	UserID   string `json:"user_id"`   // 用户ID
}
