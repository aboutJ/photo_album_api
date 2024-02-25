package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"photo_album/model"
)

type Config struct {
	Mysql Mysql
}

type Mysql struct {
	Username string
	Password string
	Url      string
	Port     string
	Database string
}

var DB *gorm.DB
var SaltPassword string = "album"

func init() {
	dataBytes, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		fmt.Println("读取配置文件失败：", err)
		return
	}

	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		fmt.Println("解析yaml文件失败：", err)
		return
	}

	err = initMysql(config.Mysql)
	if err != nil {
		fmt.Println("数据库连接失败：", err)
		return
	}
}

func initMysql(conf Mysql) error {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Url, conf.Port, conf.Database)
	var err error
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return errors.New("数据库连接失败")
	}

	// 数据迁移
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	return nil
}
