package storage

import (
	pb "progress-service/genprotos"
)

type StorageI interface {
	UserLesson() UserLessonI
	UserData() UserDataI
	Quiz() QuizI
}

type UserLessonI interface {
	Create(*pb.UserLessonCReq) (*pb.Void, error)
}

type UserDataI interface {
	GetUserData(*pb.ByID) (*pb.UserDataGRes, error)
	UpdateXP(*pb.XPUReq) (*pb.Void, error)
	UpdateDailyStreak(*pb.StreakUReq) (*pb.Void, error)
	UpdateWinningPercentage(*pb.WinningPercentageUReq) (*pb.Void, error)
	GetLeadBoard(*pb.Void) (*pb.LeadboardRes, error)
}

type QuizI interface {
	Start(*pb.TestCheckReq, *pb.ExerciseGARes) (*pb.TestResultRes, error)
}
