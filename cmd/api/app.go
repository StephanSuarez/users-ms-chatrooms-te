package api

import (
	// "fmt"

	"github.com/StephanSuarez/chat-rooms-users-ms/internal/common/config"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Env      *config.Env
	DbConn   *mongo.Database
	UsersDep *http.UsersDependencies
}

func NewApp() *App {

	app := &App{}

	app.Env = config.NewEnv()

	dbenv := &config.DbEnv{
		DbEnviroment: app.Env.DbEnviroment,
		Server:       app.Env.MongoServer,
		Username:     app.Env.MongoUsername,
		Password:     app.Env.MongoPassword,
		Cluster:      app.Env.MongoCluster,
		Dbname:       app.Env.DbName,
	}

	app.DbConn = config.GetDBInstance(dbenv)

	app.UsersDep = http.NewUsersDependencies(app.DbConn)

	http.SubcribersHandlers(app.UsersDep)

	return app
}

// func (app *App) Start() {
// 	addr := fmt.Sprintf("http://localhost:%s", app.Env.ServerAddress)
// 	log.Printf("Server is running on %s", addr)
// 	app.Router.Run(app.Env.PortServer)
// }
