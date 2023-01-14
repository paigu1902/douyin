package main

import (
	"github.com/gin-gonic/gin"
	"paigu1902/douyin/router"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run()
}
