package main

import (
	"fmt"
	"users/internal/users/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")
	r := gin.Default()
	http.Routes(r)
}
