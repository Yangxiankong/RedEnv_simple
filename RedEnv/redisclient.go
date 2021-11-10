package RedEnv

import (
	"fmt"
	redigo "github.com/garyburd/redigo/redis"
	"gopkg.in/ini.v1"
)

//var redisConn redigo.Conn

var redisPool *redigo.Pool

type redisconf struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
	PoolSize int `ini:"poolSize"`
}

func init() {
	var filepath string = "./config/redisconf.ini"
	config, err := loadRedis(filepath)
	if err != nil {
		fmt.Println("redis config error")
		return
	}
	redisPool = &redigo.Pool{
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		MaxIdle: config.PoolSize,
	}
}

func loadRedis(path string) (redisconf, error) {
	var config redisconf
	conf, err := ini.Load(path)
	if err != nil {
		fmt.Println("load redis config error")
		return config, err
	}
	err = conf.MapTo(&config)
	if err != nil {
		fmt.Println("map to redis config error")
		return config, err
	}
	fmt.Println("redisInfo : ", config)
	return config, nil
}