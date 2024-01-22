package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})

	r.Run()
}
