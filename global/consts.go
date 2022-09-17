package global

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// 初始化相关
var (
	GLOBAL_DB    *gorm.DB      //数据库
	GLOBAL_CACHE *redis.Client //缓存
	GLOBAL_LOG   *zap.Logger   // zap日志库
)

// jwt相关
const (
	TokenExpireDuration = time.Hour * 12 // token expire time

	UserId = "userId"

	TimeFormat = "2006-01-02 15:04:05"
)

var MySecret = []byte("a8x0sd.")
