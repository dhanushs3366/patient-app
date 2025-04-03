package client

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
	return container.NewHBox(card)
}

// func (c *Client) Navbar() *fyne.Container {
// 	// shouldnt clear window it is fixed a2d2ff
// 	blue := color.RGBA{R: 0xA2, G: 0xD2, B: 0xFF, A: 0xFF}
// 	homeTxt := canvas.NewText("Home", blue)
// 	homeButton := widget.NewButton(homeTxt.Text, func() {
// 		c.Window.SetContent(parseWithBorder(c.Navbar(), nil, c.About(), true))
// 	})
// 	navbarContainer := container.NewHBox(homeButton)
// 	navbarContainer.Resize(fyne.NewSize(c.About().MinSize().Width, c.Window.Canvas().Size().Height))
// 	return navbarContainer
// }

func (c *Client) Navbar() *fyne.Container {
	blue := color.RGBA{R: 0xA2, G: 0xD2, B: 0xFF, A: 0xFF}
	homeTxt := canvas.NewText("Home", blue)

	var navbarContainer *fyne.Container

	homeButton := widget.NewButton(homeTxt.Text, func() {
		// Use window content directly instead of passing navbarContainer
		c.Window.SetContent(
			container.NewBorder(navbarContainer, nil, nil, nil, c.About()),
		)
	})

	navbarContainer = container.NewHBox(homeButton)

	// Set fixed height for navbar (do NOT use dynamic sizing)
	navbarContainer.Resize(fyne.NewSize(
		c.Window.Canvas().Size().Width, // Width fills window
		50,                             // Fixed height of 50px
	))

	return navbarContainer
}
