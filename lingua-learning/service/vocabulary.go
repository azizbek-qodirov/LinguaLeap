package service

import (
	"context"
	"encoding/json"
	"fmt"
	pb "learning-service/genprotos"
	st "learning-service/storage"

	"github.com/google/uuid"
)

type VocabularyService struct {
	storage st.StorageI
	pb.UnimplementedVocabularyServiceServer
}

type VocabularyUpdateMessage struct {
	ExerciseID string `json:"exercise_id"`
	Action     string `json:"action"`
}

func (s *VocabularyService) ProcessMessage(body []byte) {
	fmt.Printf("Received message: %s\n", body)
	var msg VocabularyUpdateMessage
	err := json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Printf("Error decoding message: %v\n", err)
		return
	}

	switch msg.Action {
	case "add":
		_, err = s.AddToVocabulary(context.Background(), &pb.ByID{Id: msg.ExerciseID})
	case "delete":
		_, err = s.DeleteFromVocabulary(context.Background(), &pb.ByID{Id: msg.ExerciseID})
	}

	if err != nil {
		fmt.Printf("Error processing message: %v\n", err)
	}
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
