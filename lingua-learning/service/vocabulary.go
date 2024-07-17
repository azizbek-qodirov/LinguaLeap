package service

import (
	"context"
	"fmt"
	pb "learning-service/genprotos"
	st "learning-service/storage"

	"github.com/google/uuid"
)

type VocabularyService struct {
	storage st.StorageI
	pb.UnimplementedVocabularyServiceServer
}

func NewVocabularyService(storage st.StorageI) *VocabularyService {
	return &VocabularyService{storage: storage}
}

func (s *VocabularyService) AddToVocabulary(ctx context.Context, idReq *pb.ByID) (*pb.Void, error) {
	exercise, err := s.storage.Exercise().GetByID(idReq)
	if err != nil {
		return nil, fmt.Errorf("failed to get exercise data: %v", err)
	}
	exercise.Id = uuid.NewString()
	_, err = s.storage.Vocabulary().AddTo(idReq, exercise)
	if err != nil {
		return nil, fmt.Errorf("failed to add exercise to vocabulary: %v", err)
	}

	return &pb.Void{Success: true}, nil
}
func (s *VocabularyService) DeleteFromVocabulary(ctx context.Context, idReq *pb.ByID) (*pb.Void, error) {
	return s.storage.Vocabulary().DeleteFrom(idReq)
}

func (s *VocabularyService) GetVocabularies(ctx context.Context, req *pb.VocabulariesGAReq) (*pb.VocabulariesGARes, error) {
	return s.storage.Vocabulary().Get(req)
}
