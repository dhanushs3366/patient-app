package client

import (
	"log"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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

// use richtext for buttons

func chatHandler(c *Client) buttonHandlerFun {
	var wg sync.WaitGroup
	chatHistory := widget.NewRichText()
	chatHistory.Wrapping = fyne.TextWrapBreak

	// User input field
	userInput := widget.NewEntry()

	// implement loading progress
	chatBotMsg, err := c.chatBot.Respond("")

	initialText := &widget.TextSegment{
		Text:  chatBotMsg,
		Style: widget.RichTextStyle{},
	}

	chatHistory.Segments = append(chatHistory.Segments, initialText)

	if err != nil {
		log.Printf("Error,%s\n", err.Error())
		// implement a pop up window alerting the patient the system is down
	}

	userInput.SetPlaceHolder("Type your message here...")
	userInput.Wrapping = fyne.TextWrapBreak

	sendButton := widget.NewButton("Send", func() {
		userMsg := userInput.Text
		if userMsg != "" {

			userSegment := &widget.TextSegment{
				Text: userMsg,
				Style: widget.RichTextStyle{
					ColorName: theme.ColorNamePrimary,
				},
			}

			// Add segment to chat history
			chatHistory.Segments = append(chatHistory.Segments, userSegment)

			// Add a line break
			lineBreak := &widget.TextSegment{
				Text:  "\n",
				Style: widget.RichTextStyle{},
			}
			chatHistory.Segments = append(chatHistory.Segments, lineBreak)

			// Clear input and refresh chat display
			userInput.SetText("")
			chatHistory.Refresh()

			log.SetFlags(log.Ltime | log.Lshortfile)
			log.Printf("usr msg:%s\n", userMsg)
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
			botTextSegment := &widget.TextSegment{
				Text: botResponse,
				Style: widget.RichTextStyle{
					ColorName: theme.ColorNamePrimary,
				},
			}

			chatHistory.Segments = append(chatHistory.Segments, botTextSegment)
			chatHistory.Refresh()
		}
	})

	bookAppointment := widget.NewButton("Book", func() {
		log.Println("booking appointment")
	})

	userInputContainer := container.NewGridWithRows(2, userInput, container.NewGridWithColumns(2, sendButton, bookAppointment))
	content := container.NewBorder(nil, userInputContainer, nil, nil, chatHistory)

	handler := func() {
		c.Window.SetContent(c.Navbar(content))
	}

	return handler
}
