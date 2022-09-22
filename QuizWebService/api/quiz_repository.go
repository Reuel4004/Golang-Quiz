package api

import (
	"errors"
	"fmt"
)

//Dataset of questions
var questions = []QuizQuestion{
	{Id: 1, Question: "How many stripes are there on the US flag?", Answers: []string{"11", "15", "13", "10"}},
	{Id: 2, Question: "How many days does it take for the Earth to orbit the Sun?", Answers: []string{"365", "364", "366", "367"}},
	{Id: 3, Question: "What’s the smallest country in the world?", Answers: []string{"The Vatican", "San Marino", "Malta", "Monaco"}},
	{Id: 4, Question: "How many keys does a classic piano have?", Answers: []string{"86", "88", "84", "90"}},
}

//Dataset of answers
var answers = []QuizAnswer{
	{Id: 1, CorrectAnswer: 3},
	{Id: 2, CorrectAnswer: 1},
	{Id: 3, CorrectAnswer: 1},
	{Id: 4, CorrectAnswer: 2},
}

//Contains the history of other users. Each element is the number of correct answers posted by a user.
var correctAnswersHistory = []int{2, 4, 3, 1, 4, 3}

// AddCorrectAnswersFromUser Add the number of correct answer from a user to the history. Return the user score percentile
func AddCorrectAnswersFromUser(numberOfCorrectAnswers int) int {
	//Add correct answer to history
	correctAnswersHistory = append(correctAnswersHistory, numberOfCorrectAnswers)

	//Get the answers history lower than current number of answers
	var lowerCorrectAnswers float32
	for _, v := range correctAnswersHistory {
		if v < numberOfCorrectAnswers {
			lowerCorrectAnswers++
		}
	}
	percentile := lowerCorrectAnswers / (float32)(len(correctAnswersHistory)) * 100
	return (int)(percentile)
}

// RetrieveCorrectAnswers Given the answers from a user, calculate the number of correct answers.
func RetrieveCorrectAnswers(userAnswers []QuizAnswer) (numberOfCorrectAnswers int, err error) {
	for _, userAnswer := range userAnswers {
		correctAnswer, err := RetrieveCorrectAnswerById(userAnswer.Id)
		if err != nil {
			return -1, fmt.Errorf("there is not an answer with id %d", userAnswer.Id)
		}
		if correctAnswer == userAnswer.CorrectAnswer {
			numberOfCorrectAnswers++
		}
	}
	return numberOfCorrectAnswers, nil
}

// RetrieveQuestionById Return a question from its ID.
func RetrieveQuestionById(idBySearch int) (foundedQuestion QuizQuestion, err error) {
	// Loop through the questions, looking for the one with the given id.
	for _, a := range questions {
		if a.Id == idBySearch {
			return a, nil
		}
	}
	return foundedQuestion, errors.New("not found")
}

// RetrieveCorrectAnswerById Retrieve the correct answer from a question id
func RetrieveCorrectAnswerById(idBySearch int) (correctAnswer int, err error) {

	// Loop through the answers, looking for the one with the given id.
	for _, a := range answers {
		if a.Id == idBySearch {
			return a.CorrectAnswer, nil
		}
	}
	return -1, errors.New("not found")
}
