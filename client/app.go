package client // Ensure package matches directory name

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

type Client struct {
	App    fyne.App
	Window fyne.Window
}

func NewClient() *Client {
	return &Client{
		App:    app.New(),
		Window: app.New().NewWindow("Healthcare Assistant"),
	}
}

func (c *Client) Run() {
	c.Window.Resize(fyne.NewSize(1024, 768))
	c.Window.ShowAndRun()
}

func (c *Client) Clear() {
	c.Window.SetContent(
		container.NewHBox(),
	)
	c.Window.SetContent(c.Navbar())
}

func parseWithBorder(topOrLeft, bottomOrRight *fyne.Container, center fyne.CanvasObject, isVertical bool) *fyne.Container {
	if isVertical {
		x := container.NewBorder(topOrLeft, bottomOrRight, nil, nil, center)
		if x != nil {
			log.Println(x.Position().Components())
		} else {
			log.Println("F")
		}
		return x
	} else {
		return container.NewBorder(nil, nil, topOrLeft, bottomOrRight, center)
	}
}
