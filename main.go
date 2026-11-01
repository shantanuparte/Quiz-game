package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	amt := flag.Int("amount", 10, "Number of questions (1-50)")
	cat := flag.Int("category", 0, "Category ID (9-32, 0 for all)")
	diff := flag.String("difficulty", "", "Difficulty: easy, medium, hard")
	typ := flag.String("type", "", "Question type: multiple, boolean")

	flag.Parse()

	cli := NewClient(checkFlags(*amt, *cat, *diff, *typ))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := runQuiz(ctx, cli); err != nil {
		log.Fatal(err)
	}

}

func runQuiz(ctx context.Context, source QuestionSource) error {

	ques, err := source.GetQuestion(ctx)
	if err != nil {
		return err
	}

	game := Game{
		Questions: ques,
		Score:     0,
		Current:   0,
	}

	reader := bufio.NewReader(os.Stdin)

	total := len(game.Questions)

	fmt.Println("\n========== TRIVIA QUIZ ==========")
	fmt.Printf("Total Questions: %d\n", total)
	fmt.Println("================================")

	// Loop through all questions.
	for i, q := range game.Questions {
		game.Current = i

		correct := AskQuestion(q, i+1, total, reader)

		if correct {
			game.Score++
		}
	}

	// Display final results.
	fmt.Println("\n========== QUIZ RESULTS ==========")
	fmt.Printf("Score: %d/%d\n", game.Score, total)

	if total > 0 {
		percentage := float64(game.Score) / float64(total) * 100
		fmt.Printf("Accuracy: %.2f%%\n", percentage)
	}

	fmt.Println("==================================")

	return nil
}
