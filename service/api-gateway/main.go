package main

import (
	"github.com/gin-gonic/gin"
	"paigu1902/douyin/service/api-gateway/router"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run()
}
