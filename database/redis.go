package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
	"time"
)

var redisClient *redis.Client

func InitRedis() {
	addr := fmt.Sprintf("%v:%v", viper.GetString("redis.addr"), viper.GetString("redis.port"))
	password := viper.GetString("redis.password")
	client := redis.NewClient(&redis.Options{
		Addr:     addr, // redis地址
		Password: password, // redis没密码，没有设置，则留空
		DB:       0,  // 使用默认数据库
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	redisClient = client
	log.Println("connect redis success")
}

func SetLock(c *gin.Context, key string, value int64, expiration time.Duration) bool {
	result := redisClient.WithContext(c).SetNX(key, value, expiration)
	log.Printf("redis set lock result:%v, err:%v", result.Val(), result.Err())
	return result.Val()
}

func DelLock(c *gin.Context, key string, value int64) {
	temp, _ := redisClient.WithContext(c).Get(key).Int64()
	if temp == value {
		redisClient.WithContext(c).Del(key).Err()
	}
}