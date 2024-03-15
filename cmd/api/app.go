package api

import (
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/common/config"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env      *config.Env
	Route    *gin.Engine
	DbConn   *mongo.Database
	UsersDep *http.UsersDependencies
}

func App() Application {

	app := &Application{}

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

	app.Route = gin.Default()
	http.Routes(app.Route, app.UsersDep)
	return *app
}
