package databases

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once       sync.Once
	gormDB     *gorm.DB
	gormLogger *log.Logger
)

type MysqlConfig struct {
	DSN           string          // mysql address
	LogLevel      logger.LogLevel `json:",default=4,options=[1,2,3,4]"` // 日志等级
	SlowThreshold int             `json:",default=200"`                 // 慢sql判断条件（单位毫秒）
	LogPath       string          `json:",default=./logs/sql.log"`      // 日志文件
	Colorful      bool            `json:",optional"`                    // 彩色打印
}

// MustNewDB 单例创建db
func MustNewDB(c MysqlConfig) *gorm.DB {
	once.Do(func() {
		// 设置日志输出文件
		setLogOutput(c.LogPath)
		sqlLogger := logger.New(gormLogger, logger.Config{
			SlowThreshold:             time.Millisecond * time.Duration(c.SlowThreshold),
			Colorful:                  c.Colorful,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  c.LogLevel,
		})
		var err error
		gormDB, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       c.DSN,
			DefaultStringSize:         255,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 sqlLogger,
			NowFunc: func() time.Time {
				return time.Now().In(time.FixedZone("CST", 8*3600))
			},
			PrepareStmt:                              true,
			DisableForeignKeyConstraintWhenMigrating: true,
			QueryFields:                              true,
			CreateBatchSize:                          100,
		})
		if err != nil {
			logrus.Fatalf("[open mysql error]: %s", err.Error())
		}
		sqlDB, err := gormDB.DB()
		if err != nil {
			logrus.Fatalf("[get sql db error]: %s", err.Error())
		}
		// 设置连接最大生命周期
		sqlDB.SetConnMaxLifetime(time.Minute * 30)
		// 设置最大连接数
		sqlDB.SetMaxOpenConns(20)
		// 设置最大空闲连接数
		sqlDB.SetMaxIdleConns(5)
		// 设置最大空闲连接时间
		sqlDB.SetConnMaxIdleTime(time.Minute * 5)
		// ping
		if err = sqlDB.Ping(); err != nil {
			logrus.Fatalf("[ping mysql db error]: %s", err.Error())
		}
	})
	return gormDB
}

// mkdir 检查目录是否存在，不存在则创建
func mkdir(filePath string) {
	dir := filepath.Dir(filePath)
	_, err := os.Stat(dir)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0664)
		if err != nil {
			logrus.Fatalf("[Mkdir Error]: make dir %s error : %s", dir, err)
		}
	} else {
		logrus.Fatalf("[Stat dir Error]: stat dir %s error : %s", dir, err)
	}
}

// setLogOutput 设置输出文件
func setLogOutput(filePath string) {
	mkdir(filePath)
	// 初始化logger
	gormLogger = log.New(os.Stdout, "", log.Ltime)
	// 创建今天的日志文件
	createNewLogFile(filePath)
	//开启定时任务，每天0点替换日志
	now := time.Now()
	// 计算下一个0点
	next := now.Add(time.Hour * 24)
	next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
	t := time.NewTimer(next.Sub(now))
	go func() {
		<-t.C
		createNewLogFile(filePath)
		tc := time.NewTicker(24 * time.Hour)
		for range tc.C {
			createNewLogFile(filePath)
		}
	}()
}

// createNewLogFile 创建日志文件&设置
func createNewLogFile(filePath string) {
	dateStr := time.Now().Format("2006-01-02")
	filePath = fmt.Sprintf("%s-%s", filePath, dateStr)
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		logrus.Errorf("open %s file failed [error: %s]", filePath, err.Error())
		return
	}
	gormLogger.SetOutput(f)
}
