package RedEnv

import (
	"RedEnv_Simple/RedEnv/statuscode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GwlReceive struct {
	Uid int `json:"uid"`
}

func GwlHandler(context *gin.Context) {
	var gr GwlReceive
	if err := context.ShouldBindJSON(&gr); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": statuscode.BadRequest,
		})
		return
	}

	envList, envCnt, money, err:= getGwl(gr.Uid)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": statuscode.NoSuchUser,
		})
		return
	}
	envs := make([]gin.H, envCnt)
	for i := 0; i < envCnt; i++ {
		var flag bool
		if envList[i].Opened == 0 {
			flag = false
		} else {
			flag = true
		}
		envs[i] = gin.H{
			"envelope_id": envList[i].ID,
			"value": envList[i].Money,
			"opened": flag,
			"snatch_time": envList[i].GetTime,
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code": statuscode.Success,
		"msg": "success",
		"data" : gin.H{
			"amount" : money,
			"envelope_list": envs,
		},
	})
}