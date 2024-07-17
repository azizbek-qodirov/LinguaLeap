package service

import (
	"context"
	pb "learning-service/genprotos"
	st "learning-service/storage"
)

type ExerciseService struct {
	storage st.StorageI
	pb.UnimplementedExerciseServiceServer
}

func NewExerciseService(storage st.StorageI) *ExerciseService {
	return &ExerciseService{storage: storage}
}

func (s *ExerciseService) CreateExercise(ctx context.Context, exercise *pb.ExerciseCReqUReqForSwagger) (*pb.Void, error) {
	return s.storage.Exercise().Create(exercise)
}

func (s *ExerciseService) GetExerciseByID(ctx context.Context, idReq *pb.ByID) (*pb.ExerciseGResUReq, error) {
	return s.storage.Exercise().GetByID(idReq)
}

func (s *ExerciseService) UpdateExercise(ctx context.Context, exercise *pb.ExerciseGResUReq) (*pb.Void, error) {
	return s.storage.Exercise().Update(exercise)
}

func (s *ExerciseService) DeleteExercise(ctx context.Context, idReq *pb.ByID) (*pb.Void, error) {
	return s.storage.Exercise().Delete(idReq)
}

func (s *ExerciseService) GetAllExercises(ctx context.Context, allExercises *pb.ExerciseGAReq) (*pb.ExerciseGARes, error) {
	return s.storage.Exercise().GetAll(allExercises)
}
