package api

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

type Agent struct {
	Name         string
	Instructions string
	Client       *openai.Client
	ModelName    string
	history      []openai.ChatCompletionMessage
}

func NewAgent(name, instructions, model string) *Agent {
	authToken := os.Getenv("OPENROUTER_AI_API_KEY")
	baseURL := os.Getenv("OPENROUTER_AI_BASE_URL")
	config := openai.DefaultConfig(authToken)
	config.BaseURL = baseURL

	client := openai.NewClientWithConfig(config)
	history := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: instructions,
		},
	}
	return &Agent{
		Name:         name,
		Instructions: instructions,
		Client:       client,
		ModelName:    model,
		history:      history,
	}
}

func (a *Agent) Respond(input string) (string, error) {
	if input != "" {
		a.history = append(a.history, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})
	}
	resp, err := a.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    a.ModelName,
			Messages: a.history,
		},
	)

	if err != nil {
		log.Printf("no result from open ai, %s", err.Error())
		return "", err
	}
	a.history = append(a.history, resp.Choices[0].Message)
	return resp.Choices[0].Message.Content, nil
}

func (a *Agent) Close() {
	a.history = nil
}

func GetHealthCareAssistantAgent() *Agent {
	instructions := `
		Your role is health care assistant robot, and you go by personal health care assistant
		First Greet the patient, and introduce yourselves
		You are a health care assistant robot, you are assigned as a doctor's assistant.
		Your role is to act like a personal assistant and help patient with normal enquiries with possible precautions they can take to avoid their health issue.
		You should also provide assistance via word of mouth. Patient might come for a mental health issue or they might have a phyical bruise but you should be supportive and helpful.
		They might give you what their symptoms are if it is a sympton that can lead to other health issues you should inform the patient and give them proper precautions they should take.
		Try to get the more info from the patient but ask easy questions dont give too many options it will overwhelm them
		only ask a single line or two line questions to get info

		after you feel like you narrowed their problem and given enough instructions tell the patient to book apointment using our chatbot feature if they are still unsure about it
	`
	modelName := "mistralai/mistral-7b-instruct:free"

	return NewAgent("Health Care Agent", instructions, modelName)
}
