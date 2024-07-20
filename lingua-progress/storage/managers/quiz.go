package managers

import (
	"database/sql"
	"fmt"

	pb "progress-service/genprotos"
)

type QuizManager struct {
	Conn *sql.DB
}

func NewQuizManager(conn *sql.DB) *QuizManager {
	return &QuizManager{Conn: conn}
}

func (m *QuizManager) Start(quizzes *pb.TestCheckReq, exercises *pb.ExerciseGARes) (*pb.TestResultRes, error) {
	exerciseMap := make(map[string]*pb.ExerciseGResUReq)
	for _, exercise := range exercises.Exercises {
		exerciseMap[exercise.Id] = exercise
	}

	var correctAnswersCount int
	var totalQuestions int

	for _, quiz := range quizzes.Requests {
		totalQuestions++
		exercise, exists := exerciseMap[quiz.ExerciseId]
		if !exists {
			continue
		}
		if exercise.CorrectAnswer == quiz.CorrectAnswer {
			correctAnswersCount++
		}
	}

	xpGiven := correctAnswersCount * 3
	feedback := fmt.Sprintf("You got %d out of %d questions correct!", correctAnswersCount, totalQuestions)

	return &pb.TestResultRes{
		TestsCount:          int32(totalQuestions),
		CorrectAnswersCount: int32(correctAnswersCount),
		XpGiven:             int32(xpGiven),
		Feedback:            feedback,
	}, nil
}
