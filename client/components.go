package client

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (c *Client) Navbar(content *fyne.Container) *fyne.Container {
	blue := color.RGBA{R: 0xA2, G: 0xD2, B: 0xFF, A: 0xFF}
	homeTxt := canvas.NewText("Home", blue)

	homeButton := widget.NewButton(homeTxt.Text, func() {
		c.Window.SetContent(c.Navbar(content))
	})

	// Create navbar with spacer for right alignment of exit button
	navbarContainer := container.NewHBox(homeButton, layout.NewSpacer(), c.streamBtn(), layout.NewSpacer(), c.exitBtn())

	paddedContent := padContainer(content, true, true)
	return container.NewBorder(navbarContainer, nil, nil, nil, paddedContent)
}

func (c *Client) About() *fyne.Container {

	pink := color.RGBA{R: 255, G: 192, B: 203, A: 255}
	cardTitle := canvas.NewText("What is HealthCare Assistant", pink)
	cardContent := widget.NewLabel("Hi im your personal robot.")

	card := widget.NewCard(
		cardTitle.Text,
		"Subtitle",
		cardContent,
	)
	iconImage := canvas.NewImageFromResource(c.App.Icon())
	card.SetImage(iconImage)
	aboutContainer := container.NewHBox(card)
	return aboutContainer
}

func (c *Client) Stream() *fyne.Container {
	// should have a container to scan your face

	return nil

}
