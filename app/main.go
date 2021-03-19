package main

import (
	"go-starter/app/cmd"
	_ "go-starter/docs"
)

// @title go-starter restful API
// @version 1.0.0
// @description go-starter
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://localhost:8080
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cmd.Execute()
}
