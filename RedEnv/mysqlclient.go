package RedEnv

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type mysqlconf struct {
	User string `ini:"user"`
	Password string `ini:"password"`
	Host string `ini:"host"`
	Port string `ini:"port"`
	Db string `ini:"db"`
	Param string `ini:"param"`
}

func init() {
	var filepath string = "./config/mysqlconf.ini"
	config, err := loadmysql(filepath)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", config.User, config.Password, config.Host, config.Port, config.Db, config.Param)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect error")
		return
	}
}

func loadmysql(path string) (mysqlconf, error) {
	var config mysqlconf
	conf, err := ini.Load(path)
	if err != nil {
		fmt.Println("load mysql config file error!")
		return config, err
	}
	conf.BlockMode =false
	err = conf.MapTo(&config)
	if err != nil {
		fmt.Println("map to mysql config error")
	}
	fmt.Println("mysqlInfo : ", config)
	return config, nil
}
