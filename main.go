package main

import "github.com/gin-gonic/gin"
import "RedEnv_Simple/RedEnv"

func main() {
	router := gin.Default()

	RedEnv.LoadRedEnv(router)

	router.Run(":9499")
}