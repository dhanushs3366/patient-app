package client

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (c *Client) DoctorLogin() *fyne.Container {
	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Enter username")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Enter password")

	loginButton := widget.NewButton("Login", func() {
		// Your login button logic here
		log.Println("Login button pressed with username:", usernameEntry.Text)
		// Add authentication logic
	})

	loginForm := container.New(
		layout.NewGridLayoutWithRows(3),
		layout.NewSpacer(),
		container.New(
			layout.NewGridLayoutWithColumns(3),
			layout.NewSpacer(),
			container.NewVBox(
				usernameEntry,
				passwordEntry,
				loginButton,
			),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)

	// Set button importance to change its color

	// Create layout with form and button
	formContainer := container.NewVBox(
		loginForm,
		loginForm,
	)

	return formContainer
}
