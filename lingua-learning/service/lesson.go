package service

import (
	"context"
	pb "learning-service/genprotos"
	st "learning-service/storage"
)

type LessonService struct {
	storage st.StorageI
	pb.UnimplementedLessonServiceServer
}

func NewLessonService(storage st.StorageI) *LessonService {
	return &LessonService{storage: storage}
}

func (s *LessonService) CreateLesson(ctx context.Context, category *pb.LessonCReqGRes) (*pb.Void, error) {
	return s.storage.Lesson().Create(category)
}

func (s *LessonService) GetByID(ctx context.Context, idReq *pb.ByID) (*pb.LessonCReqGRes, error) {
	return s.storage.Lesson().GetByID(idReq)
}

func (s *LessonService) Update(ctx context.Context, category *pb.LessonUReq) (*pb.Void, error) {
	return s.storage.Lesson().Update(category)
}

func (s *LessonService) Delete(ctx context.Context, idReq *pb.ByID) (*pb.Void, error) {
	return s.storage.Lesson().Delete(idReq)
}

func (s *LessonService) GetAll(ctx context.Context, allCategories *pb.LessonGAReq) (*pb.LessonGARes, error) {
	return s.storage.Lesson().GetAll(allCategories)
}
