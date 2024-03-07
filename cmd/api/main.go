package main

import (
	"users/internal/common/config"
	"users/internal/users/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	env      *config.Env
	route    *gin.Engine
	dbConn   *mongo.Database
	usersDep *http.UsersDependencies
}

func main() {

	app := App{}

	app.env = config.NewEnv()

	dbenv := &config.DbEnv{
		DbEnviroment: app.env.DbEnviroment,
		Server:       app.env.MongoServer,
		Username:     app.env.MongoUsername,
		Password:     app.env.MongoPassword,
		Cluster:      app.env.MongoCluster,
		Dbname:       app.env.DbName,
	}

	app.dbConn = config.GetDBInstance(dbenv)

	app.usersDep = http.NewUsersDependencies(app.dbConn)

	app.route = gin.Default()
	http.Routes(app.route, app.usersDep)
}
