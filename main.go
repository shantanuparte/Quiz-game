
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	amt := flag.Int("amount", 10, "Number of questions (1-50)")
	cat := flag.Int("category", 0, "Category ID (9-32, 0 for all)")
	diff := flag.String("difficulty", "", "Difficulty: easy, medium, hard")
	typ := flag.String("type", "", "Question type: multiple, boolean")

	flag.Parse()

	if *amt < 1 || *amt > 50 {
		fmt.Println("Amount must be between 1 and 50")
		os.Exit(1)
	}

	if *cat != 0 && (*cat < 9 || *cat > 32) {
		fmt.Println("Category must be between 9 and 32")
		os.Exit(1)
	}

	if *diff != "" && *diff != "easy" && *diff != "medium" && *diff != "hard" {
		fmt.Println("Invalid difficulty")
		os.Exit(1)
	}

	if *typ != "" && *typ != "multiple" && *typ != "boolean" {
		fmt.Println("Invalid question type")
		os.Exit(1)
	}

	params := url.Values{}
	params.Set("amount", fmt.Sprintf("%d", *amt))

	if *cat != 0 {
		params.Set("category", fmt.Sprintf("%d", *cat))
	}

	if *diff != "" {
		params.Set("difficulty", *diff)
	}

	if *typ != "" {
		params.Set("type", *typ)
	}

	baseURL := "https://opentdb.com/api.php"
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(apiURL)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("HTTP error:", resp.Status)
		return
	}

	var data Response

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("JSON decoding failed:", err)
		return
	}

	out, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("JSON encoding failed:", err)
		return
	}

	fmt.Println(string(out))
}
