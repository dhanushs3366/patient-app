package client // Ensure package matches directory name

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/dhanushs3366/patient-app/api"
	"github.com/dhanushs3366/patient-app/db"
	"github.com/dhanushs3366/patient-app/db/models"
)

type Client struct {
	App            fyne.App
	Window         fyne.Window
	chatBot        *api.Agent
	store          *db.Store
	loggedInDoctor *models.Doctor
	config         *WindowConfig
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
		App:            app.New(),
		Window:         app.New().NewWindow("Healthcare Assistant"),
		config:         nil,
		chatBot:        api.GetHealthCareAssistantAgent(),
		store:          store,
		loggedInDoctor: nil,
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
