package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up Our App")
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
