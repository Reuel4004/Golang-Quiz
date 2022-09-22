package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var baseAddress string = "http://localhost:8080"

func GetQuestions() ([]QuizQuestion, error) {
	result, err := executeGet("questions")
	if err != nil {
		return nil, err
	}

	var questions []QuizQuestion
	jsonErr := json.Unmarshal(result, &questions)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return questions, nil
}

func PostAnswers(answers []QuizAnswer) (quizResult QuizResult, err error) {
	body, _ := json.Marshal(answers)
	result, err := executePost("correctAnswers", body)
	if err != nil {
		return quizResult, err
	}
	jsonErr := json.Unmarshal(result, &quizResult)
	if jsonErr != nil {
		return quizResult, jsonErr
	}
	return quizResult, nil
}

func executeGet(endpoint string) (result []byte, apiError error) {
	url := fmt.Sprintf("%s/%s", baseAddress, endpoint)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	result, err = readResponse(response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func executePost(endpoint string, body []byte) (result []byte, apiError error) {
	url := fmt.Sprintf("%s/%s", baseAddress, endpoint)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	result, err = readResponse(response)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func readResponse(response *http.Response) (result []byte, apiError error) {
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
