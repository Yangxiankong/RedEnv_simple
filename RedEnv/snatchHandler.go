package RedEnv

import (
	"RedEnv_Simple/RedEnv/statuscode"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

type UserID struct {
	Uid int `json:"uid"`
}

func SnatchHandler(context *gin.Context) {
	N := 5		//最多抢几次
	p := 0.8	//抢到的概率

	//测试
	var usr UserID

	if err := context.ShouldBindJSON(&usr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": statuscode.BadRequest,	//请求格式有误无法解析
			"msg": "request format error",
		})
		return
	}

	fmt.Println("uid", usr.Uid)
	//拿到usr的count
	cnt, err := getSnatch(usr.Uid)
	fmt.Println("cnt", cnt)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": statuscode.NoSuchUser,	//没有这名用户
			"msg": "no such user",
		})
		return
	}

	//通过count返回红包信息
	if cnt >= N {
		context.JSON(http.StatusOK, gin.H{
			"code": statuscode.CantGetMoreEnv,	//获取红包达到上限
			"msg": "cant get more red envelope",
		})
		return
	}

	rnum := rand.Float64()
	if rnum <= p {
		currEnvid++
		money := rand.Intn(4950) + 50
		env := Env{currEnvid, usr.Uid, money, 0, time.Now().Unix()}
		usrcnt := UserCount{usr.Uid, cnt + 1}
		go updateSnatch(env, usrcnt)
		context.JSON(http.StatusOK, gin.H{
			"code": statuscode.Success,
			"msg": "success",
			"data": gin.H{
				"envelope_id": currEnvid,
				"max_count": N,
				"cur_count": cnt + 1,
			},
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"code": statuscode.Thankyou,	//没抢到
			"msg": "thank you",
		})
	}

}