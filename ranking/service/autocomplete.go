package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"ranking/domain"
)

type AutoCompleteSvc interface {
	CompleteSearchBar(ctx context.Context, keyword string) ([]domain.Movie, error)
}

type autocompleteSvc struct {
}

func NewAutoCompleteSvc() AutoCompleteSvc {
	return &autocompleteSvc{}
}

func (a autocompleteSvc) CompleteSearchBar(ctx context.Context, keyword string) ([]domain.Movie, error) {
	ip := os.Args[1]
	url := "http://" + ip + ":9200/movie_catalog/_search"

	// create query content
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  keyword,
				"fields": []string{"*"},
			},
		},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		log.Printf("Error encoding query to JSON: %s", err)
		return nil, err
	}

	// create autocomplete HTTP request to elastic search
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(queryJSON))
	if err != nil {
		log.Printf("Error creating HTTP request: %s", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error executing HTTP request: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return nil, err
	}

	// Check for not OK status codes
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from server: %s", resp.Status)
	}

	// Parse and display the results
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error parsing JSON response: %s", err)
	}

	hits := int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	if hits <= 0 {
		return []domain.Movie{}, nil
	}

	res := make([]domain.Movie, hits)
	for idx, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var movie domain.Movie
		source := hit.(map[string]interface{})["_source"]
		sourceJSON, err := json.Marshal(source)
		if err != nil {
			log.Printf("Error marshaling source to JSON: %s", err)
			return res, nil
		}
		if err := json.Unmarshal(sourceJSON, &movie); err != nil {
			log.Printf("Error unmarshaling JSON to Movie: %s", err)
			return res, nil
		}
		res[idx] = movie
	}
	return res, nil
}
