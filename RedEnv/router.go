package RedEnv

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

//var redisConn redis.Conn
var currEnvid int

func LoadRedEnv(router *gin.Engine) {
	/*	net := "tcp"
		addr := "192.168.14.128:6379"*/
	var env Env
	result := db.Last(&env)
	if result.RowsAffected == 0 {
		currEnvid = 0
	} else {
		currEnvid = env.ID
	}
	/*	var err error
		redisConn, err = redis.Dial(net, addr)
		if err != nil {
			fmt.Println("connect to redis failed")
		}*/
	rand.Seed(time.Now().UnixNano())
	//defer redisConn.Close()

	router.POST("snatch", SnatchHandler)
	router.POST("open", OpenHandler)
	router.POST("get_wallet_list", GwlHandler)
}
