package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	initializeRoute(router)
	router.Run(":8080")
}
