package handlers

import (
	pb "gateway-service/genprotos"

	"google.golang.org/grpc"
)

type HTTPHandler struct {
	Lesson     pb.LessonServiceClient
	Exercise   pb.ExerciseServiceClient
	Vocabulary pb.VocabularyServiceClient
	UserLesson pb.UserLessonServiceClient
	UserData   pb.UserDataServiceClient
}

func NewHandler(connL, connP *grpc.ClientConn) *HTTPHandler {
	return &HTTPHandler{
		Lesson:     pb.NewLessonServiceClient(connL),
		Exercise:   pb.NewExerciseServiceClient(connL),
		Vocabulary: pb.NewVocabularyServiceClient(connL),
		UserLesson: pb.NewUserLessonServiceClient(connP),
		UserData:   pb.NewUserDataServiceClient(connP),
	}
}
