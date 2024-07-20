package service

import (
	"context"
	"errors"
	pb "progress-service/genprotos"
	"progress-service/storage"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

type QuizService struct {
	storage               storage.StorageI
	LessonServiceClient   pb.LessonServiceClient
	ExerciseServiceClient pb.ExerciseServiceClient
	pb.UnimplementedQuizServiceServer
	RabbitMQ *amqp.Connection
}

func NewQuizService(storage storage.StorageI, learningConn *grpc.ClientConn, rabbitMQ *amqp.Connection) *QuizService {
	lessonClient := pb.NewLessonServiceClient(learningConn)
	exerciseClient := pb.NewExerciseServiceClient(learningConn)
	return &QuizService{storage: storage, LessonServiceClient: lessonClient, ExerciseServiceClient: exerciseClient, RabbitMQ: rabbitMQ}
}

func (s *QuizService) StartTest(ctx context.Context, req *pb.TestCheckReq) (*pb.TestResultRes, error) {
	lesson, err := s.LessonServiceClient.GetLessonByID(ctx, &pb.ByID{Id: req.LessonId})
	if err != nil {
		return nil, err
	}
	if lesson == nil {
		return nil, errors.New("lesson not found")
	}
	exercises, err := s.ExerciseServiceClient.GetAllExercises(ctx, &pb.ExerciseGAReq{LessonId: req.LessonId})
	if err != nil {
		return nil, err
	}
	res, err := s.storage.Quiz().Start(req, exercises)
	if err != nil {
		return nil, err
	}

	_, err = s.storage.UserData().UpdateWinningPercentage(&pb.WinningPercentageUReq{UserId: req.UserId, Percentage: float32(res.CorrectAnswersCount) * 100 / float32(res.TestsCount)})
	if err != nil {
		return nil, err
	}
	_, err = s.storage.UserData().UpdateXP(&pb.XPUReq{UserId: req.UserId, Xp: int64(res.XpGiven)})
	if err != nil {
		return nil, err
	}
	_, err = s.storage.UserData().UpdateDailyStreak(&pb.StreakUReq{UserId: req.UserId, DailyStreak: 1})
	if err != nil {
		return nil, err
	}
	_, err = s.storage.UserLesson().Create(
		&pb.UserLessonCReq{
			Id:       uuid.NewString(),
			UserId:   req.UserId,
			LessonId: req.LessonId,
		},
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}
