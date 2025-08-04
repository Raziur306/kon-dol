package worker

import (
	"fmt"
	"log"
	"os"

	"github.com/Raziur306/kon-dol/internal/fetcher"
	"github.com/Raziur306/kon-dol/internal/gpt"
	"github.com/Raziur306/kon-dol/internal/utils"
)

func PullNewsList() {
	uri := os.Getenv("LATEST_NEWS_BN_API")
	baseUrl := os.Getenv("BASE_URL")

	if uri == "" {
		log.Fatal("LATEST_NEWS_BN_API environment variable is not set")
		return
	}

	if baseUrl == "" {
		log.Fatal("BASE_URL environment variable is not set")
		return
	}

	newsList, err := fetcher.PullNewsList(uri)
	if err != nil {
		log.Fatalf("Error pulling news list: %v", err)
		return
	}

	//extract and pull each news item
	fmt.Println("Found ", len(newsList.Items), "news items")

	// Sequentially process each item
	for i, item := range newsList.Items {

		newsResponse, err, lastUpdatedAt := fetcher.FetchSingleNewsFullContext(item.Slug)
		if err != nil {
			log.Printf("Error fetching full context for item %d: %v", i, err)
			continue
		}
		mainContext := utils.ExtractSummaryContent(newsResponse.Data.Story)

		if mainContext == "" {
			log.Printf("No main context found for item %d", i)
			continue
		}

		//check if the item is violated or not
		if utils.IsPotentialPoliticalViolence(mainContext, item.Tags) {
			log.Printf("Item %d is not violated, skipping...", i)
			continue
		}

		// Process with GPT
		fmt.Printf("Processing item %d with GPT...\n", i)
		gpt.ProcessWithGPT(item.URL, mainContext, lastUpdatedAt, item.TrackId)

	}

}
