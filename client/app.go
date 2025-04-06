package client // Ensure package matches directory name

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/dhanushs3366/patient-app/api"
	"github.com/dhanushs3366/patient-app/db"
)

type Client struct {
	App     fyne.App
	Window  fyne.Window
	chatBot *api.Agent
	store   *db.Store
	config  *WindowConfig
}

type WindowConfig struct {
	Navbar dimensions
}

type dimensions struct {
	Height float64
	Width  float64
}

func NewClient() (*Client, error) {
	store, err := db.New()

	if err != nil {
		return nil, err
	}

	if err := store.Init(); err != nil {
		return nil, err
	}

	return &Client{
		App:     app.New(),
		Window:  app.New().NewWindow("Healthcare Assistant"),
		config:  nil,
		chatBot: api.GetHealthCareAssistantAgent(),
		store:   store,
	}, nil
}

func (c *Client) Run() {
	c.Window.Resize(fyne.NewSize(1024, 768))
	c.Window.ShowAndRun()
}

func (c *Client) Clear() {
	c.Window.SetContent(
		container.NewHBox(),
	)
	c.Window.SetContent(c.Navbar(c.About()))
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

func padContainer(content *fyne.Container, hPad, vPad bool) *fyne.Container {
	var top, bottom, left, right fyne.CanvasObject
	if vPad {
		top = layout.NewSpacer()
		bottom = layout.NewSpacer()
	}
	if hPad {
		left = layout.NewSpacer()
		right = layout.NewSpacer()
	}

	return container.NewPadded(
		container.NewVBox(
			top,
			container.NewHBox(
				left,
				content,
				right,
			),
			bottom,
		),
	)

}
