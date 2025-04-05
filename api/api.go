package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type HTTPReq string

const (
	GET    HTTPReq = "GET"
	POST   HTTPReq = "POST"
	DELETE HTTPReq = "DELETE"
	UPDATE HTTPReq = "UPDATE"
)

type APIClient struct {
	BaseURL string
	APIKey  string
	Client  *http.Client
}

func NewAPIClient() *APIClient {
	URL := os.Getenv("BASE_URL")
	APIKey := os.Getenv("OPENAI_API_KEY")
	client := http.Client{Timeout: 20 * time.Second}
	return &APIClient{
		BaseURL: URL,
		APIKey:  APIKey,
		Client:  &client,
	}
}

func (c *APIClient) MakeRequest(method, endpoint string, payload interface{}, headers map[string]string) ([]byte, error) {
	url := c.BaseURL + endpoint

	var req *http.Request

	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Couldnt marshal payload, %s\n", err.Error())
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			log.Printf("Couldnt create new req, %s\n", err.Error())
			return nil, err
		}
	} else {
		// idk whats happening here if i declare err like this
		// req,err:=...
		// it throws a linting error
		// if i do it like this
		// req,err =
		// it throws an error too
		req, _ = http.NewRequest(method, url, nil)

	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New("error: " + resp.Status)
	}

	return io.ReadAll(resp.Body)
}
