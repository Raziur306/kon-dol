package gpt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Raziur306/kon-dol/internal/db"
	"github.com/Raziur306/kon-dol/internal/model"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func ProcessWithGPT(link string, mainContext string, lastUpdatedAt int64, id string) {
	if mainContext == "" {
		log.Println("Main context is empty, skipping GPT processing")
		return
	}

	//check already exist or not
	_, _, defaultCollection := db.ConnectDB()
	if defaultCollection == nil {
		log.Fatal("Failed to connect to MongoDB")
		return
	}

	existingNews := defaultCollection.FindOne(context.Background(), model.Incident{TrackId: id})

	if existingNews != nil {
		log.Printf("News with TrackId %s already exists, skipping GPT processing", id)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
    "name": "name of the party like// "বিএনপি", "আওয়ামী লীগ", "জাতীয় পার্টি", "জামায়াত", "জাতীয় সমাজতান্ত্রিক দল", "নগর উন্নয়ন পার্টি", "ইসলামী আন্দোলন", "বাসদ", "ন্যাপ", "কমিউনিস্ট পার্টি", "নতুন ধারা", "গণফোরাম", "নাগরিক ঐক্য", "বাম মোর্চা", "লেবার পার্টি", "আমরা জনগণ পার্টি", "নিপ", "এনসিপি", "জাতীয় নাগরিক পার্টি", ",
    "type": "political | unknown"
  },
  "date": "...",
  "short_desc": "...",
  "status:"..."// can be violation or not violation
}

The text of the article is in Bengali. set it to an empty string.
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

	var result model.GPTResponse
	err = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal GPT response: %v", err)
	}

	if result.Date == "" {
		result.Date = time.Unix(lastUpdatedAt, 0).Format("2006-01-02")
	}

	//format result to match the Incident model
	incident := model.Incident{
		Title:     result.Title,
		Location:  result.Location,
		Party:     result.Party.Name,
		Date:      result.Date,
		Source:    link,
		ShortDesc: result.ShortDesc,
		TrackId:   id,
	}

	_, err = defaultCollection.InsertOne(ctx, incident)
	if err != nil {
		log.Fatalf("Failed to insert GPT result into MongoDB: %v", err)
		return
	}

	log.Println("✅ GPT result successfully stored in MongoDB")

}
