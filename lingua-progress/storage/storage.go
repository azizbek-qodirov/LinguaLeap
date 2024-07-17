package storage

import (
	pb "progress-service/genprotos"
)

type StorageI interface {
	UserLesson() UserLessonI
	UserData() UserDataI
}

type UserLessonI interface {
	Create(*pb.UserLessonCReq) (*pb.Void, error)
}

type UserDataI interface {
	GetUserData(*pb.ByID) (*pb.UserDataGRes, error)
	UpdateXP(*pb.XPUReq) (*pb.Void, error)
	UpdateDailyStreak(*pb.StreakUReq) (*pb.Void, error)
	UpdatePlayedGamesCount(*pb.PlayedGamesCountUReq) (*pb.Void, error)
	UpdateWinningPercentage(*pb.WinningPercentageUReq) (*pb.Void, error)
}
