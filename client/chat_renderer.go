package client

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
* it should take chat response from agent.Respond() and append it into a container of some transparent bg
 */

func renderChat(responseTxt string) *fyne.Container {

	// rich text to represent and style the text

	// a transparent container

	chatLabel := widget.NewLabel(responseTxt)
	chatLabel.Wrapping = fyne.TextWrapWord

	chatBox := canvas.NewRectangle(color.RGBA{R: 50, G: 150, B: 200, A: 150})
	chatBox.CornerRadius = 10

	chatBubble := container.NewStack(chatBox, chatLabel)

	return chatBubble
}
