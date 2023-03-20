package gptlib

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/PullRequestInc/go-gpt3"
)

type Client interface {
	SendRequest(RequestData) (string, error)
}

type RequestData struct {
	// required
	Prompt string

	// optional
	MaxTokens int    // by default: 1000
	Role      string // by default: user
	Timeout   time.Duration
	UserID    string
}

func NewChatGPT(openAIAPIToken string) Client {
	return &chatGPT{
		conn: gpt3.NewClient(openAIAPIToken),
	}
}

type chatGPT struct {
	conn gpt3.Client
}

func (c *chatGPT) SendRequest(data RequestData) (string, error) {
	if data.Role == "" {
		data.Role = "user"
	}
	if data.MaxTokens == 0 {
		data.MaxTokens = 1000
	}

	var ctx context.Context
	var ctxCancel func()
	if data.Timeout == 0 {
		ctx = context.Background()
	} else {
		ctx, ctxCancel = context.WithTimeout(context.Background(), data.Timeout)
	}
	defer ctxCancel()

	response, err := c.conn.ChatCompletion(
		ctx,
		gpt3.ChatCompletionRequest{
			Messages: []gpt3.ChatCompletionRequestMessage{
				{
					Role:    data.Role,
					Content: data.Prompt,
				},
			},
			Temperature:      0.6,
			MaxTokens:        data.MaxTokens,
			TopP:             1,
			N:                1,
			FrequencyPenalty: 1,
			PresencePenalty:  1,
			User:             data.UserID,
		},
	)
	if err != nil {
		return "", fmt.Errorf("send chat completion request: %w", err)
	}

	jsonByes, _ := json.Marshal(response)
	fmt.Println(string(jsonByes))

	dataArray := []string{}
	for _, data := range response.Choices {
		if data.Message.Content != "" {
			dataArray = append(dataArray, data.Message.Content)
		}
	}

	result := strings.Join(dataArray, "\n")
	return strings.TrimLeft(result, "\n"), nil
}
