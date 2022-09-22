## Quiz
The task is to build a super simple quiz with a few questions and a few alternatives for each question. With one correct answer per question.

## Preferred Stack
Backend - Golang
Database - Just in-memory , so no database

## Preferred Components
REST API or gRPC
CLI that talks with the API, preferably using https://github.com/spf13/cobra ( as cli framework )

## User stories/Use cases
- User should be presented questions with a number of answers
- User should be able to select just one answer per question
- User should be able to answer all the questions and then post his/her answers and get back how many correct answers there had and be displayed to the user.
- User should see how good he/she rated compared to others that have taken the quiz, "You scored higher than 60% of all quizzers" 

## How to Run
Web Service : Run QuizWebService\main.go. It will receive requests on localhost:8080. 
Quiz Console: Run QuizConsole\main.go from cmd, using "quiz" as argument.
