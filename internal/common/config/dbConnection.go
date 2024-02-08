package config

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbEnv struct {
	Server   string
	Username string
	Password string
	Cluster  string
	Dbname   string
}

var dbInstance *mongo.Database
var lock = &sync.Mutex{}

func mongoConnection(dbEnv *DbEnv) {

	mongoURI := fmt.Sprintf("%s://%s:%s@%s/", dbEnv.Server, dbEnv.Username, dbEnv.Password, dbEnv.Cluster)

	opts := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	dbInstance = client.Database(dbEnv.Dbname)
}

func GetDBInstance(dbEnv *DbEnv) *mongo.Database {

	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		mongoConnection(dbEnv)
	}

	return dbInstance
}
