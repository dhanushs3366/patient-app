package client

import (
	"log"

	"fyne.io/fyne/v2/widget"
)

func (c *Client) exitBtn() *widget.Button {
	exit := widget.NewButton("Exit", func() {
		log.Println("Exiting")
		c.App.Quit()
	})

	return exit
}

func (c *Client) streamBtn() *widget.Button {
	stream := widget.NewButton("Stream", func() {
		log.Println("Streaming")
	})

	return stream
}

// func(c *Client) refreshBtn() *widget.Button{
// 	refresh:=widget.NewButton("Refresh", func() {
// 		c.Window.Canvas().Refresh()
// 	})
// }
