package client

import (
	"log"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type buttonHandlerFun func()

func (c *Client) exitBtn() *widget.Button {
	exit := widget.NewButton("Exit", func() {
		log.Println("Exiting")
		c.App.Quit()
	})

	return exit
}

func (c *Client) chatBtn() *widget.Button {
	chat := widget.NewButton("Chat with me", chatHandler(c))

	return chat
}

func (c *Client) streamBtn() *widget.Button {
	stream := widget.NewButton("Stream", func() {
		log.Println("Streaming")
	})

	return stream
}

func (c *Client) Login() *widget.Button {
	login := widget.NewButton("Login", func() {
		c.Window.SetContent(c.Navbar(c.DoctorLogin()))
	})

	return login
}

// use richtext for buttons

func chatHandler(c *Client) buttonHandlerFun {
	var wg sync.WaitGroup
	// add scrollable feature to chat bubbles
	chatBubbles := []fyne.CanvasObject{}
	var chatBubble *fyne.Container
	// User input field
	userInput := widget.NewEntry()

	// implement loading progress
	chatBotMsg, err := c.chatBot.Respond("")
	if err != nil {
		promptWindow("Error", "AI bot is down", &c.Window)
		c.Window.SetContent(c.Navbar(c.About()))
	}

	chatBubble = renderChat(chatBotMsg, BOT)
	chatBubbles = append(chatBubbles, chatBubble)

	// a global contianer that contains all the chat bubbles
	// used to refresh the chatbubbles
	chatBox := container.NewVBox(chatBubbles...)

	userInput.SetPlaceHolder("Type your message here...")
	userInput.Wrapping = fyne.TextWrapBreak

	sendButton := widget.NewButton("Send", func() {
		userMsg := userInput.Text
		if userMsg != "" {

			chatBubble = renderChat(userMsg, PATIENT)
			chatBubbles = append(chatBubbles, chatBubble)
			chatBox.Add(chatBubble)
			// Clear input and refresh chat display
			userInput.SetText("")
			chatBox.Refresh()

			var botResponse string
			wg.Add(1)
			go func() {
				defer wg.Done()
				botResponse, err = c.chatBot.Respond(userMsg)
				log.Printf("bot response:%s\n", botResponse)
				if err != nil {
					log.Printf("Error getting response from chatbot:%s\n", err.Error())
					// implement a prompt to alert the patient if system is down
				}
			}()

			wg.Wait()

			chatBubble = renderChat(botResponse, BOT)
			chatBubbles = append(chatBubbles, chatBubble)
			chatBox.Add(chatBubble)
			chatBox.Refresh()
		}
	})

	bookAppointment := widget.NewButton("Book", func() {
		c.Window.SetContent(c.Navbar(c.BookDoctor()))
	})

	clearChat := widget.NewButton("Clear", func() {
		c.chatBot.Close()
		chatBubbles = []fyne.CanvasObject{}
		chatBox.Objects = nil
		respTxt, err := c.chatBot.Respond("")
		if err != nil {
			promptWindow("Error", err.Error(), &c.Window)
		}

		chatBubbles = append(chatBubbles, renderChat(respTxt, BOT))
		chatBox.Add(chatBubbles[len(chatBubbles)-1])
		chatBox.Refresh()
	})

	userInputContainer := container.NewGridWithRows(2, userInput, container.NewGridWithColumns(3, sendButton, bookAppointment, clearChat))
	// add a toolbar for filtering doctors
	chatBox = container.NewVBox(chatBubbles...)
	content := container.NewBorder(nil, userInputContainer, nil, nil, container.NewVScroll(chatBox))

	handler := func() {
		c.Window.SetContent(c.Navbar(content))
	}

	return handler
}
