package main

import (
	"fmt"

	"github.com/StephanSuarez/chat-rooms-users-ms/cmd/api"
)

func main() {
	app := api.App()
	fmt.Println(app)
}
