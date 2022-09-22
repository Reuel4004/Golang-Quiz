package api

// QuizQuestion Represent a question from the quiz.
type QuizQuestion struct {
	Id       int      `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

// QuizAnswer Represent the answer to a quiz question
type QuizAnswer struct {
	Id            int `json:"id"`
	CorrectAnswer int `json:"correct_answer"`
}
