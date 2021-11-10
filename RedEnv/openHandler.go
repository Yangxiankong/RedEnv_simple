package RedEnv

import (
	"RedEnv_Simple/RedEnv/statuscode"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserEnv struct {
	Uid int `json:"uid"`
	Envelope_id int `json:"envelope_id"`
}

func OpenHandler(context *gin.Context) {
	var ue UserEnv
	if err := context.ShouldBindJSON(&ue); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"code": statuscode.BadRequest,
			"msg": "request format error",
		})
		return
	}
	fmt.Println("ue : ", ue)
	envVal, money, err := getOpen(ue.Envelope_id, ue.Uid)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code": statuscode.NoSuchRec,
			"msg": "no such user and red envelope record",
		})
		return
	}
	go updateOpen(ue.Envelope_id, ue.Uid, money + envVal)
	context.JSON(http.StatusOK, gin.H{
		"code": statuscode.Success,
		"msg": "success",
		"data": gin.H{
			"value": envVal,
		},
	})
}