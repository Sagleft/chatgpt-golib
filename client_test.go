package gptlib

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	msg := `{"id":"chatcmpl-6wCAKTIeQhXVCJJmDNwmSVbics6s4","object":"chat.completion","created":1679327808,"model":"gpt-3.5-turbo-0301","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\n\nAs an AI language model, I do not have a definite answer to this question as it is still unknown how many stars are there in our galaxy. However, according to the latest estimates by astronomers and scientists, there could be anywhere between 100 billion to 400 billion stars in the Milky Way Galaxy."}}],"usage":{"prompt_tokens":16,"completion_tokens":62,"total_tokens":78}}`

	var response *gpt3.ChatCompletionResponse
	require.NoError(t, json.Unmarshal([]byte(msg), &response))

	text := getResponseText(response)
	assert.NotEqual(t, "", text)

	fmt.Println(text)
}
