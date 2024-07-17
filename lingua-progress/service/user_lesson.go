package service

import (
	"context"
	pb "progress-service/genprotos"
	st "progress-service/storage"

	"github.com/google/uuid"
)

type UserLessonService struct {
	storage st.StorageI
	pb.UnimplementedUserLessonServiceServer
}

func NewUserLessonService(storage st.StorageI) *UserLessonService {
	return &UserLessonService{storage: storage}
}

func (s *UserLessonService) Create(ctx context.Context, req *pb.UserLessonCReq) (*pb.Void, error) {
	req.Id = uuid.NewString()
	return s.storage.UserLesson().Create(req)
}
