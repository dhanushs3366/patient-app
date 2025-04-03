package main

import "github.com/dhanushs3366/patient-app/client"

func main() {
	c := client.NewClient()
	c.Window.SetContent(c.Navbar())

	c.Window.ShowAndRun()

	// myApp := app.New()
	// myWindow := myApp.NewWindow("TabContainer Widget")

	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
	// 	container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	// )

	// //tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	// tabs.SetTabLocation(container.TabLocationLeading)

	// myWindow.SetContent(tabs)
	// myWindow.ShowAndRun()
}
