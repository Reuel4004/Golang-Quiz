package cmd

import (
	"QuizConsole/api"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// quizCmd represents the quiz command
var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Quiz Started! Select one answer for each question.")

		//Get questions from api source.
		questions, err := api.GetQuestions()
		if err != nil {
			fatalError(err)
		}

		//Start quiz and get user answers
		answers := startQuiz(questions)

		//Post answers and get score
		quizResult, err := api.PostAnswers(answers)
		if err != nil {
			fatalError(err)
		}
		printResult(quizResult)
	},
}

// Start the quiz. Show questions and get answers
func startQuiz(questions []api.QuizQuestion) []api.QuizAnswer {
	answers := make([]api.QuizAnswer, 0, len(questions))
	for _, v := range questions {
		printQuestion(v)
		answer := api.QuizAnswer{
			Id:            v.Id,
			CorrectAnswer: getAnswer(len(v.Answers)),
		}
		answers = append(answers, answer)
	}
	return answers
}

// Get user answer for a question
func getAnswer(numberOfAnswers int) int {

	//Read selection from answer
	var typedAnswer string
	for {
		//Scan user answer
		_, err := fmt.Scanln(&typedAnswer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		//Check inf answer is number
		id, err := strconv.Atoi(typedAnswer)
		if err != nil {
			fmt.Println("Only numbers are accepted")
			continue
		}

		//Check if number is included in answers range
		if id > 0 && id <= numberOfAnswers {
			return id
		} else {
			fmt.Println("Given number is not an answer")
		}
	}

}

// Print the score
func printResult(result api.QuizResult) {
	fmt.Printf("Number of correct answers: %d\n", result.NumberOfCorrectAnswers)
	fmt.Println(result.Message)
}

//Print a question and its answer
func printQuestion(question api.QuizQuestion) {
	fmt.Printf("Question number: %d\n", question.Id)
	fmt.Println(question.Question)

	//Print questions one by one
	for i, v := range question.Answers {
		fmt.Printf("%d) %s\n", i+1, v)
	}
}

func fatalError(err error) {
	fmt.Printf("Error %s", err)
	os.Exit(1)
}

func init() {
	rootCmd.AddCommand(quizCmd)
}
