package http

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.POST("/users", AddUser)

	r.Run()
}
