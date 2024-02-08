package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/welcome", GetUsers)
	r.POST("/users", AddUser)

	r.Run()
}
