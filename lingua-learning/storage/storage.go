package storage

import (
	pb "learning-service/genprotos"
)

type StorageI interface {
	Lesson() LessonI
	Vocabulary() VocabularyI
	Exercise() ExerciseI
}

type LessonI interface {
	Create(*pb.LessonCReqGRes) (*pb.Void, error)
	GetByID(*pb.ByID) (*pb.LessonCReqGRes, error)
	Update(*pb.LessonUReq) (*pb.Void, error)
	Delete(*pb.ByID) (*pb.Void, error)
	GetAll(*pb.LessonGAReq) (*pb.LessonGARes, error)
}

type VocabularyI interface {
	AddTo(*pb.ByID, *pb.ExerciseGResUReq) (*pb.Void, error)
	DeleteFrom(*pb.ByID) (*pb.Void, error)
	Get(*pb.VocabulariesGAReq) (*pb.VocabulariesGARes, error)
}

type ExerciseI interface {
	Create(*pb.ExerciseCReqUReqForSwagger) (*pb.Void, error)
	GetByID(*pb.ByID) (*pb.ExerciseGResUReq, error)
	Update(*pb.ExerciseGResUReq) (*pb.Void, error)
	Delete(*pb.ByID) (*pb.Void, error)
	GetAll(*pb.ExerciseGAReq) (*pb.ExerciseGARes, error)
}
