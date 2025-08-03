package gpt

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func ProcessWithGPT(link string, mainContext string) {
	token := os.Getenv("GPT_TOKEN")
	baseUrl := os.Getenv("GPT_BASE_URL")
	if token == "" {
		log.Fatal("GPT Token Not Found")
		return
	}
	if baseUrl == "" {
		log.Fatal("GPT Base URL Not Found")
		return
	}
	client := openai.NewClient(
		option.WithBaseURL(baseUrl),
		option.WithAPIKey(token),
	)

	prompt := fmt.Sprintf(`You are an intelligent news incident extractor. Based on the following article, extract the details and return in **valid JSON** that matches the following structure: {
  "title": "...",
  "location": "write incident location here",
  "district": "write incident district name here",
  "party": {
    "name": "...",
    "type": "political | unknown"
  },
  "date": "...",
  "short_desc": "...",
}

The text of the article is in Bengali. Translate where needed. If any field is not found, set it to an empty string.
Article:
%s
`, mainContext)

	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT4_1,
	})

	if err != nil {
		log.Fatalf("Error creating chat completion: %v", err)
		return
	}

	fmt.Println("Message", chatCompletion.Choices[0].Message.Content)

}
