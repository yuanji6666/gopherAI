package redis

import (
	"fmt"

	"github.com/yuanji6666/gopherAI/config"
)

func GenerateCaptcha(email string) string{
	return fmt.Sprintf(config.DefaultRedisKeyConfig.CaptchaPrefix, email)
}