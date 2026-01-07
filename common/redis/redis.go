package redis

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yuanji6666/gopherAI/config"
)

var Rdb *redis.Client

var ctx = context.Background()

func InitRedis(){
	conf := config.GetConfig()
	host := conf.RedisHost
	port := conf.RedisPort
	password := conf.RedisPassword
	db := conf.RedisDB

	addr := host + ":" + strconv.Itoa(port)

	Rdb = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: db,
	})
}

func SetCaptchaForEmail(email, captcha string) error{
	key := GenerateCaptcha(email)
	expire := 2*time.Minute
	return  Rdb.Set(ctx, key, captcha, expire).Err()
}

func CheckCaptchaForEmail(email, userInputCaptcha string) (bool , error){
	key := GenerateCaptcha(email)

	storedCaptcha, err := Rdb.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return false, nil
		}

		return false, nil
	}

	if strings.EqualFold(storedCaptcha, userInputCaptcha) {
		Rdb.Del(ctx, key)
		return true, nil
	}

	return false, nil
}



