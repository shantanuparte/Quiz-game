package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Question struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type Response struct {
	ResponseCode int        `json:"response_code"`
	Results      []Question `json:"results"`
}

type OpClient struct {
	BaseUrl    string
	HttpClient *http.Client
}

type QuestionSource interface {
	GetQuestion(ctx context.Context) ([]Question, error)
}

func NewClient(brul string) *OpClient {
	return &OpClient{
		BaseUrl: brul,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func checkFlags(amount, cato int, diff, typ string) string {
	if amount < 1 || amount > 50 {
		fmt.Println("Amount must be between 1 and 50")
		os.Exit(1)
	}

	if cato != 0 && (cato < 9 || cato > 32) {
		fmt.Println("Category must be between ")
		os.Exit(1)
	}

	if diff != "" && diff != "easy" && diff != "medium" && diff != "hard" {
		fmt.Println("Invalid difficulty")
		os.Exit(1)
	}

	if typ != "" && typ != "multiple" && typ != "boolean" {
		fmt.Println("Invalid question type")
		os.Exit(1)
	}

	params := url.Values{}
	params.Set("amount", fmt.Sprintf("%d", amount))

	if cato != 0 {
		params.Set("category", fmt.Sprintf("%d", cato))
	}

	if diff != "" {
		params.Set("difficulty", diff)
	}

	if typ != "" {
		params.Set("type", typ)
	}

	baseUrl := "https://opentdb.com/api.php"
	apiUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())

	return apiUrl

}

func (c *OpClient) GetQuestion(ctx context.Context) ([]Question, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseUrl, nil)

	if err != nil {
		return nil, fmt.Errorf("Creating request: %w", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Sending request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api returned http status: %s", resp.Status)
	}
	
	var data Response

	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("Decoding response: %w", err)
	}

	if data.ResponseCode != 0 {
		return nil, fmt.Errorf("Open TDB returned respose Code %d", data.ResponseCode)
	}

	for i := range data.Results {
		q := &data.Results[i]

		q.Question = html.UnescapeString(q.Question)
		q.CorrectAnswer = html.UnescapeString(q.CorrectAnswer)

		for j := range q.IncorrectAnswers {
			q.IncorrectAnswers[j] = html.UnescapeString(q.IncorrectAnswers[j])
		}

	}

	return data.Results, nil

}
