package function

import (
	"fmt"
	"net/http"

	"github.com/StephanSuarez/chat-rooms-users-ms/cmd/api"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	app := api.App()
	fmt.Println(app)
	functions.HTTP("UsersMS", usersMS)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func usersMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Users-ms")
}
