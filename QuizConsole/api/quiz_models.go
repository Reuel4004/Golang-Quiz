package api

type QuizQuestion struct {
	Id       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

type QuizAnswer struct {
	Id            int `json:"id"`
	CorrectAnswer int `json:"correct_answer"`
}

type QuizResult struct {
	NumberOfCorrectAnswers int    `json:"correct_answers"`
	Message                string `json:"message"`
}
