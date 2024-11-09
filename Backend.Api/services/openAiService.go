package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type requestBody struct {
	Model       string    `json:"model"`
	Messages    []message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

type responseBody struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

type OpenAIClient struct {
	APIKey string
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{APIKey: apiKey}
}

func (client *OpenAIClient) SendMessage(systemPrompt, userPrompt string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	reqBody := requestBody{
		Model: "gpt-4o-mini",
		Messages: []message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		Temperature: 0.7,
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+client.APIKey)
	req.Header.Set("Content-Type", "application/json")

	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var resBody responseBody
	if err := json.Unmarshal(respBody, &resBody); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	if len(resBody.Choices) > 0 {
		return resBody.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response from openAi")
}
