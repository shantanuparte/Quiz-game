package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Game struct {
	Questions []Question
	Score     int
	Current   int
}

func suffleAsnwer(q Question) ([]string, int) {

	answers := append([]string{}, q.IncorrectAnswers...)
	answers = append(answers, q.CorrectAnswer)

	rand.Shuffle(len(answers), func(i, j int) {
		answers[i], answers[j] = answers[j], answers[i]
	})

	correct := 0

	for i, answers := range answers {
		if answers == q.CorrectAnswer {
			correct = i
			break
		}
	}

	return answers, correct

}

func AskQuestion(q Question, number, total int, reader *bufio.Reader) bool {

	ans, cor := suffleAsnwer(q)

	fmt.Printf("\nQuestion %d/%d\n", number, total)
	fmt.Println(q.Question)

	for i, answer := range ans {
		fmt.Printf("[%d] %s\n", i+1, answer)
	}

	fmt.Printf("\nYour answer: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)

	if err != nil || choice < 1 || choice > len(ans) {
		fmt.Println("Invalid choice")
		return false
	}

	if choice-1 == cor {
		fmt.Println("Correct!")
		return true
	}

	fmt.Println("Incorret!")
	fmt.Println("Correct answer: ", ans[cor])

	return false

}

func ShowCategoryIDs() {
	fmt.Println("\nAvailable Category IDs (9-32):")

	categories := map[int]string{
		9:  "General Knowledge",
		10: "Entertainment: Books",
		11: "Entertainment: Film",
		12: "Entertainment: Music",
		13: "Entertainment: Musicals & Theatres",
		14: "Entertainment: Television",
		15: "Entertainment: Video Games",
		16: "Entertainment: Board Games",
		17: "Science & Nature",
		18: "Science: Computers",
		19: "Science: Mathematics",
		20: "Mythology",
		21: "Sports",
		22: "Geography",
		23: "History",
		24: "Politics",
		25: "Art",
		26: "Celebrities",
		27: "Animals",
		28: "Vehicles",
		29: "Entertainment: Comics",
		30: "Science: Gadgets",
		31: "Entertainment: Japanese Anime & Manga",
		32: "Entertainment: Cartoon & Animations",
	}

	for id := 9; id <= 32; id++ {
		fmt.Printf("  %2d : %s\n", id, categories[id])
	}
}
