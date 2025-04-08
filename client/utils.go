package client

import (
	"fmt"
	"image/color"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MyCustomTheme struct {
	baseTheme fyne.Theme
}

func (t *MyCustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// Override colors as needed
	if name == theme.ColorNameButton {
		return color.RGBA{R: 0x2C, G: 0x7A, B: 0xB8, A: 255}
	}
	return t.baseTheme.Color(name, variant)
}

func (t *MyCustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return t.baseTheme.Icon(name)
}

func (t *MyCustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return t.baseTheme.Font(style)
}

func (t *MyCustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return t.baseTheme.Size(name)
}

func (c *Client) showCustomPrompt(title, message string) *dialog.Dialog {
	infoDialog := dialog.NewInformation(title, message, c.Window)
	return &infoDialog
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

func promptWindow(title, info string, window *fyne.Window) {
	dialog.ShowInformation(title, info, *window)
}

// if valid prompt with valid msg
// if invalid prompt with your own invalid msgs
func checkValidForms(entires map[*widget.Label]*widget.Entry) (bool, string) {
	isValid := true
	invalidLabels := []string{}
	for label, entry := range entires {
		// just check if they are not empty for now
		if entry.Text == "" {
			isValid = false
			invalidLabels = append(invalidLabels, label.Text)
		}
	}

	if len(invalidLabels) > 0 {
		var invalidStr strings.Builder
		invalidStr.WriteString("Invalid data provided for: ")
		for i, label := range invalidLabels {
			if i == len(invalidLabels)-1 {
				invalidStr.WriteString(strings.TrimSuffix(label, ":"))
				break
			}
			invalidStr.WriteString(fmt.Sprintf("%s,", strings.TrimSuffix(label, ":")))
		}

		return isValid, invalidStr.String()

	}

	return isValid, ""
}
