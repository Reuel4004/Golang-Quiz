package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetQuestions Get all the quiz questions
func GetQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

// GetQuestionById Get a question from a given ID
func GetQuestionById(c *gin.Context) {
	//Get the 'ID' field from request
	idParam := c.Params.ByName("id")
	if idParam == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing 'id' field"})
		return
	}

	// Convert the ID field to a number
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be a valid integer number"})
		return
	}

	// Get question by id and return it
	foundedQuestion, err := RetrieveQuestionById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "question not found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, foundedQuestion)
	}
}

// GetNumberOfQuestions Get total number of questions available trough API
func GetNumberOfQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"count": len(questions)})
}

// PostCorrectAnswers Post quiz answers from a user. Return number of correct answers and a message of comparison with other users.
func PostCorrectAnswers(c *gin.Context) {
	var postedAnswers []QuizAnswer

	// Deserialize the json posted by the user to a collection of QuizAnswer
	if err := c.BindJSON(&postedAnswers); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "data not in a correct format"})
		return
	}

	//Get number of correct answers
	numberOfCorrectAnswers, err := RetrieveCorrectAnswers(postedAnswers)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//Calculate the percentile to know the performance of the user compared to the others
	percentile := AddCorrectAnswersFromUser(numberOfCorrectAnswers)

	// Send back number of correct answers and the performance compared to other users
	c.IndentedJSON(http.StatusOK, gin.H{"correct_answers": numberOfCorrectAnswers, "message": fmt.Sprintf("You scored higher than %d %% of all quizzers", percentile)})
}
