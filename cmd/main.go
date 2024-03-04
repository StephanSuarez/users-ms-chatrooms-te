package main

import (
	"fmt"
	"users/internal/common/config"
	"users/internal/users/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	env    *config.Env
	route  *gin.Engine
	dbConn *mongo.Database
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
	fmt.Print("------")
	fmt.Print(dbenv)

	app.dbConn = config.GetDBInstance(dbenv)

	app.route = gin.Default()
	http.Routes(app.route)
}
