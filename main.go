package main

import (
	// "encoding/json"
	"flag"
	// "net/http"
	// "time"
)

func main() {
	amt := flag.Int("amount", 10, "Number of questions (1-50)")
	cat := flag.Int("category", 0, "Category ID (9-32, 0 for all)")
	diff := flag.String("difficulty", "", "Difficulty: easy, medium, hard")
	typ := flag.String("type", "", "Question type: multiple, boolean")

	flag.Parse()

	var cli *OpClient

	cli = NewClient(checkFlags(*amt, *cat, *diff, *typ))

	// client := &http.Client{
	// 	Timeout: 10 * time.Second,
	// }
	// resp, err := client.Get(apiURL)
	// if err != nil {
	// 	fmt.Println("Request failed:", err)
	// 	return
	// }
	// defer resp.Body.Close()
	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Println("HTTP error:", resp.Status)
	// 	return
	// }
	// var data Response
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	fmt.Println("JSON decoding failed:", err)
	// 	return
	// }
	// out, err := json.MarshalIndent(data, "", "    ")
	// if err != nil {
	// 	fmt.Println("JSON encoding failed:", err)
	// 	return
	// }
	// fmt.Println(string(out))
}
