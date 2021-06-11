package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/Kira1108/production-ready-api/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up Our App")
	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest Api Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting rest api")
		fmt.Println(err)
	}
}
