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
		c.Window.SetContent(c.Navbar(container.NewCenter(c.About())))
	})

	// Create navbar with spacer for right alignment of exit button
	test_btn := widget.NewButton("Test", func() {
		c.Window.SetContent(c.Navbar(container.NewCenter(widget.NewLabel("TESTING HERE :)"))))
	})
	navbarContainer := container.New(layout.NewCustomPaddedHBoxLayout(10), homeButton, c.chatBtn(), c.Login(), c.exitBtn(), c.streamRoomBtn(), test_btn)

	// paddedContent := padContainer(content, true, true)
	return container.NewBorder(navbarContainer, nil, nil, nil, content)
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
