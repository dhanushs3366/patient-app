package client

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
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
