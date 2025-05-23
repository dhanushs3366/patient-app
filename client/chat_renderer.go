package client

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

/*
* it should take chat response from agent.Respond() and append it into a container of some transparent bg
 */

type ROLE string

const (
	PATIENT ROLE = "PATIENT"
	BOT     ROLE = "BOT"
)

func renderChat(responseTxt string, user ROLE) *fyne.Container {
	botRes := loadImageAsIcon("assets/chatbot.svg")
	personRes := loadImageAsIcon("assets/person.svg")

	// Create image objects
	botImg := canvas.NewImageFromResource(botRes)
	personImg := canvas.NewImageFromResource(personRes)

	botImg.FillMode = canvas.ImageFillContain
	personImg.FillMode = canvas.ImageFillContain

	botImg.SetMinSize(fyne.NewSize(48, 48))
	personImg.SetMinSize(fyne.NewSize(48, 48))

	botIcon := container.NewStack(botImg)
	personIcon := container.NewStack(personImg)

	responseTxt = removeNewLines(responseTxt)
	lines := wrapText(responseTxt, float32(1700))
	lineLabel := widget.NewLabel("")
	for i, line := range lines {
		lineLabel.Text += line
		if i < len(lines)-1 {
			lineLabel.Text += "\n"
		}
	}

	// Chat bubble styling
	chatBox := canvas.NewRectangle(color.RGBA{R: 50, G: 150, B: 200, A: 150})
	chatBox.CornerRadius = 10

	bubble := container.NewPadded(container.NewStack(chatBox, lineLabel))

	// Final layout
	if user == PATIENT {
		return container.NewHBox(layout.NewSpacer(), bubble, personIcon)
	} else {
		return container.NewHBox(botIcon, bubble, layout.NewSpacer())
	}
}

func wrapText(text string, maxWidth float32) []string {
	var lines []string

	paragraphs := strings.Split(text, "\n")

	for _, para := range paragraphs {
		if para == "" {
			lines = append(lines, "")
			continue
		}

		words := strings.Fields(para)
		currentLine := ""

		for _, word := range words {
			testLine := currentLine
			if currentLine != "" {
				testLine += " "
			}
			testLine += word

			testLabel := widget.NewLabel(testLine)
			if testLabel.MinSize().Width > maxWidth && currentLine != "" {
				lines = append(lines, currentLine)
				currentLine = word // start new line
			} else {
				currentLine = testLine
			}
		}

		if currentLine != "" {
			lines = append(lines, currentLine)
		}
	}

	return lines
}

func removeNewLines(txt string) string {
	var result []rune
	var prev rune

	for _, ch := range txt {
		if ch == '\n' && prev == '\n' {
			continue // skip repeated newline
		}
		result = append(result, ch)
		prev = ch
	}

	return string(result)
}
