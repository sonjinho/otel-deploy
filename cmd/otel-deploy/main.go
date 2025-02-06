package main

import "github.com/sonjinho/otel-deploy-backend/internal/app"

func main() {
	app := app.NewApplication()
	app.Run()
}
