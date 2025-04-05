package api

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func TestModel() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	agent := GetHealthCareAssistantAgent()
	fmt.Println("Hello there, how are you doing?")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("You: ")
		inp, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %s", err.Error())
			continue
		}
		inp = strings.TrimSpace(inp)

		log.Println("Waiting for response...")
		out, err := agent.Respond(inp)
		log.Println("Response received.")

		if err != nil {
			log.Printf("Error from agent: %s\n", err.Error())
			continue
		}

		fmt.Println("Assistant:", out)
	}
}
