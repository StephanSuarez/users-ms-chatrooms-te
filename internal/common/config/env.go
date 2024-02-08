package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
aaaa
type Env struct {
	MongoServer   string
	MongoUsername string
	MongoPassword string
	MongoCluster  string
	DbName        string
}

func NewEnv() *Env {
	env := Env{}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	env.MongoServer = os.Getenv("MONGO_SERVER")
	env.MongoUsername = os.Getenv("MONGO_USERNAME")
	env.MongoPassword = os.Getenv("MONGO_PASSWORD")
	env.MongoCluster = os.Getenv("MONGO_CLUSTER")
	env.DbName = os.Getenv("DB_NAME")

	return &env
}
