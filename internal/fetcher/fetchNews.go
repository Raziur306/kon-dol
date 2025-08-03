package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Raziur306/kon-dol/internal/model"
)

func PullNewsList(uri string) (model.NewsResponse, error) {

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal("Error fetching news list:", err)
		return model.NewsResponse{}, err
	}
	defer resp.Body.Close()

	//read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
		return model.NewsResponse{}, err
	}

	var newsResp model.NewsResponse
	if err := json.Unmarshal(body, &newsResp); err != nil {
		log.Fatalf("Failed to parse JSON: %v\n", err)
		return model.NewsResponse{}, err
	}

	return newsResp, nil

}

func FetchSingleNewsFullContext(uri string) (model.SingleNewsResponse, error) {
	baseUrl := os.Getenv("SINGLE_NEWS_API")

	if baseUrl == "" {
		log.Fatal("SINGLE_NEWS_API environment variable is not set")
		return model.SingleNewsResponse{}, fmt.Errorf("SINGLE_NEWS_API environment variable is not set")
	}

	resp, err := http.Get(baseUrl + uri)
	if err != nil {
		log.Fatal("Error fetching news full context:", err)
		return model.SingleNewsResponse{}, err
	}
	defer resp.Body.Close()

	var newsResp model.SingleNewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&newsResp); err != nil {
		log.Fatal("Error decoding JSON:", err)
		return model.SingleNewsResponse{}, err
	}
	return newsResp, nil
}
