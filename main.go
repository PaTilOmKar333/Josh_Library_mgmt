package main

import (
	"net/http"
	"project/app"
	"project/config"
	"project/handler"
)

func main() {
	// intialise config
	// coonnect with database
	// depedecies initialconfig.Load()
	config.Load()
	app.Init()
	defer app.Close()

	handler.InitDependancies()

	http.ListenAndServe(":8080", handler.Router())

}
