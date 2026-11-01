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
