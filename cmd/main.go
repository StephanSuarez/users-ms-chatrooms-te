package main

import (
	"users/internal/users/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	http.Routes(r)
}
