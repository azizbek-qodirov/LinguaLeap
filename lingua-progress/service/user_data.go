package service

import (
	"context"
	pb "progress-service/genprotos"
	"progress-service/storage"
)

type UserDataService struct {
	storage storage.StorageI
	pb.UnimplementedUserDataServiceServer
}

func NewUserDataService(storage storage.StorageI) *UserDataService {
	return &UserDataService{storage: storage}
}

func (s *UserDataService) GetUserData(ctx context.Context, req *pb.ByID) (*pb.UserDataGRes, error) {
	return s.storage.UserData().GetUserData(req)
}

func (s *UserDataService) UpdateXP(ctx context.Context, req *pb.XPUReq) (*pb.Void, error) {
	return s.storage.UserData().UpdateXP(req)
}

func (s *UserDataService) UpdateDailyStreak(ctx context.Context, req *pb.StreakUReq) (*pb.Void, error) {
	return s.storage.UserData().UpdateDailyStreak(req)
}

func (s *UserDataService) UpdatePlayedGamesCount(ctx context.Context, req *pb.PlayedGamesCountUReq) (*pb.Void, error) {
	return s.storage.UserData().UpdatePlayedGamesCount(req)
}

func (s *UserDataService) UpdateWinningPercentage(ctx context.Context, req *pb.WinningPercentageUReq) (*pb.Void, error) {
	return s.storage.UserData().UpdateWinningPercentage(req)
}
