package mysql

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yuanji6666/gopherAI/config"
	"github.com/yuanji6666/gopherAI/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitMysql() error {
	host := config.GetConfig().MysqlHost
	port := config.GetConfig().MysqlPort
	user := config.GetConfig().MysqlUser
	password := config.GetConfig().MysqlPassword
	DBname := config.GetConfig().MysqlDBName
	charSet := config.GetConfig().MysqlCharset
	

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", user, password, host, port, DBname, charSet)

	var log logger.Interface

	if gin.Mode() == "debug" {
		log = logger.Default.LogMode(logger.Info)
	}else{
		log = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}),&gorm.Config{
		Logger: log,
	})

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	
	DB = db

	return migration()

}

func migration() error {
	return DB.AutoMigrate(
		new(model.User),
	)
}

func InsertUser(user *model.User) (*model.User, error ){
	err := DB.Create(user).Error
	return user, err
}

func GetUserByUsername(username string) (*model.User, error){
	var user model.User
	err := DB.Where("username=?", username).First(&user).Error
	return &user, err
}