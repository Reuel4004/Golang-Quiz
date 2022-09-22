package main

import (
	"QuizWebService/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/questions", api.GetQuestions)
	router.GET("/questions/:id", api.GetQuestionById)
	router.GET("/questionsNumber", api.GetNumberOfQuestions)
	router.POST("correctAnswers", api.PostCorrectAnswers)

	router.Run("localhost:8080")
}
