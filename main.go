package main

import (
	"strconv"

	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/common/redis"
	"github.com/yuanji6666/gopherAI/config"
	"github.com/yuanji6666/gopherAI/router"
)

func main(){
	config.InitConfig()
	mysql.InitMysql()
	redis.InitRedis()

	r := router.InitRouter()
	r.Run(config.GetConfig().Host+":"+strconv.Itoa(config.GetConfig().Port))
}
