package main

import (
	"log"

	"github.com/alielmi98/golang-todo-list-api/config"
	"github.com/alielmi98/golang-todo-list-api/constants"
	"github.com/alielmi98/golang-todo-list-api/data/cache"
	"github.com/alielmi98/golang-todo-list-api/data/db"
)

func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatalf("caller:%s  Level:%s  Msg:%s", constants.Redis, constants.Startup, err.Error())
	}
	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatalf("caller:%s  Level:%s  Msg:%s", constants.Postgres, constants.Startup, err.Error())
	}
}
