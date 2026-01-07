package user

import (
	"github.com/yuanji6666/gopherAI/common/mysql"
	"github.com/yuanji6666/gopherAI/model"
	"github.com/yuanji6666/gopherAI/utils"
)

func IsExistUser(username string) (bool, *model.User){
	user, err := mysql.GetUserByUsername(username)
	if err == nil {
		return true, user
	}else{
		return false, nil
	}
}

func Register(username ,email, password string) (user *model.User, ok bool){
	user, err := mysql.InsertUser(&model.User{
		Username: username,
		Name: username,
		Email: email,
		Password: utils.MD5(password),
	})
	if err != nil {
		return nil, false
	}
	return user, true
}