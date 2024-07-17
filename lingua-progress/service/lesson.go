package service

import (
	"context"
	pb "progress-service/genprotos"
	st "progress-service/storage"
)

type UserLessonService struct {
	storage st.StorageI
	pb.UnimplementedUserLessonServiceServer
}

func NewUserLessonService(storage st.StorageI) *UserLessonService {
	return &UserLessonService{storage: storage}
}

func (s *UserLessonService) CreateUserLesson(ctx context.Context, category *pb.UserLessonCReq) (*pb.Void, error) {
	return s.storage.UserLesson().Create(category)
}
