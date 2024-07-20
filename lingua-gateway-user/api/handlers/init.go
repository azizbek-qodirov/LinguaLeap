package handlers

import (
	pb "gateway-service/genprotos"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Lesson     pb.LessonServiceClient
	Exercise   pb.ExerciseServiceClient
	Vocabulary pb.VocabularyServiceClient
	UserLesson pb.UserLessonServiceClient
	UserData   pb.UserDataServiceClient
	Quiz       pb.QuizServiceClient
	RabbitMQ   *amqp.Channel
}

func NewHandler(connL, connP *grpc.ClientConn, ch *amqp.Channel) *HTTPHandler {
	return &HTTPHandler{
		Lesson:     pb.NewLessonServiceClient(connL),
		Exercise:   pb.NewExerciseServiceClient(connL),
		Vocabulary: pb.NewVocabularyServiceClient(connL),
		UserLesson: pb.NewUserLessonServiceClient(connP),
		UserData:   pb.NewUserDataServiceClient(connP),
		Quiz:       pb.NewQuizServiceClient(connP),
		RabbitMQ:   ch,
	}
}
