package main

import (
	"github.com/dhanushs3366/patient-app/client"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}
	c := client.NewClient()
	c.Window.SetContent(c.Navbar(c.About()))
	c.Window.SetFullScreen(true)
	c.Window.ShowAndRun()

}
