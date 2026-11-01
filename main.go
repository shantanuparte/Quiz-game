package main

import (
	"context"
	"flag"
	"fmt"
	"log"
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

	for i, q := range ques {
		fmt.Printf("\n%d. %s\n", i+1, q.Question)
		fmt.Printf("Category: %s | Difficulty: %s\n", q.Category, q.Difficulty)

		fmt.Println("Correct answer: ", q.CorrectAnswer)
		fmt.Println("Incorrect answer: ", q.IncorrectAnswers)

	}

	return nil

}
