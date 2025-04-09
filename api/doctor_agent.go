package api

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

type Agent struct {
	Name              string
	Instructions      string
	Client            *openai.Client
	DeployedModelName string
	history           []openai.ChatCompletionMessage
}

func NewAgent(name, instructions, deploymentName string) *Agent {
	apiKey := os.Getenv("AZURE_AI_API_KEY")
	endpoint := os.Getenv("AZURE_AI_BASE_URL")
	apiVersion := os.Getenv("AZURE_AI_API_VERSION")

	if apiKey == "" || endpoint == "" || apiVersion == "" {
		log.Fatal("Missing Azure OpenAI configuration (API key, base URL, or version)")
	}

	config := openai.DefaultAzureConfig(apiKey, endpoint)
	config.APIVersion = apiVersion
	config.AzureModelMapperFunc = func(model string) string {
		return deploymentName
	}

	client := openai.NewClientWithConfig(config)

	return &Agent{
		Name:              name,
		Instructions:      instructions,
		Client:            client,
		DeployedModelName: deploymentName, // deploymentName used here
		history:           []openai.ChatCompletionMessage{},
	}
}

func (a *Agent) Respond(input string) (string, error) {
	if input != "" {
		a.history = append(a.history, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})
	} else {
		//  if string is empty
		//  either its initialisation of the agent or its history is cleared
		a.history = append(a.history, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: a.Instructions,
		})
	}
	resp, err := a.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    a.DeployedModelName,
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
	a.history = []openai.ChatCompletionMessage{}
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
		only ask a double line or triple line question

		after you feel like you narrowed their problem and given enough instructions tell the patient to book apointment using our chatbot feature if they are still unsure about it~
	`
	// modelName := "mistralai/mistral-7b-instruct:free"
	modelName := "gpt-4o"

	return NewAgent("Health Care Agent", instructions, modelName)
}
