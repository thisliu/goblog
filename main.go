package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/logger"
	"net/http"
)

func main() {
	database.Initialize()
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe("127.0.0.1:3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
