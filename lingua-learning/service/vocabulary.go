package service

import (
	"context"
	pb "learning-service/genprotos"
	st "learning-service/storage"
)

type VocabularyService struct {
	storage st.StorageI
	pb.UnimplementedVocabularyServiceServer
}

func NewVocabularyService(storage st.StorageI) *VocabularyService {
	return &VocabularyService{storage: storage}
}

func (s *VocabularyService) AddTo(ctx context.Context, idReq *pb.ByID) (*pb.Void, error) {
	return s.storage.Vocabulary().AddTo(idReq)
}

func (s *VocabularyService) DeleteFrom(ctx context.Context, idReq *pb.ByID) (*pb.Void, error) {
	return s.storage.Vocabulary().DeleteFrom(idReq)
}

func (s *VocabularyService) Get(ctx context.Context, req *pb.VocabulariesGAReq) (*pb.VocabulariesGARes, error) {
	return s.storage.Vocabulary().Get(req)
}
