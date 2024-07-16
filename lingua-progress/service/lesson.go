package service

import (
	"context"
	pb "progress-service/genprotos"
	st "progress-service/storage"
)

type LessonService struct {
	storage st.Storage
	pb.UnimplementedLessonServiceServer
}

func NewLessonService(storage *st.Storage) *LessonService {
	return &LessonService{storage: *storage}
}

func (s *LessonService) CreateLesson(ctx context.Context, category *pb.LessonCReqGRes) (*pb.Void, error) {
	return nil, nil
}

func (s *LessonService) GetByID(ctx context.Context, idReq *pb.ByID) (*pb.LessonCReqGRes, error) {
	return nil, nil
}

func (s *LessonService) Update(ctx context.Context, category *pb.LessonUReq) (*pb.Void, error) {
	return nil, nil
}

func (s *LessonService) Delete(ctx context.Context, idReq *pb.ByID) (*pb.Void, error) {
	return nil, nil
}

func (s *LessonService) GetAll(ctx context.Context, allCategories *pb.LessonGAReq) (*pb.LessonGARes, error) {
	return nil, nil
}
