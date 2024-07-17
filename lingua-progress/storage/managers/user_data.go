package managers

import (
	"context"
	"time"

	pb "progress-service/genprotos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDataManager struct {
	Collection *mongo.Collection
}

func NewUserDataManager(client *mongo.Client, dbName, collectionName string) *UserDataManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserDataManager{Collection: collection}
}

func (m *UserDataManager) GetUserData(req *pb.ByID) (*pb.UserDataGRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result pb.UserDataGRes
	filter := bson.M{"user_id": req.Id}
	err := m.Collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (m *UserDataManager) UpdateXP(req *pb.XPUReq) (*pb.Void, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": req.UserId}
	update := bson.M{"$set": bson.M{"xp": req.Xp}}
	_, err := m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{Success: true}, nil
}

func (m *UserDataManager) UpdateDailyStreak(req *pb.StreakUReq) (*pb.Void, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": req.UserId}
	update := bson.M{"$set": bson.M{"daily_streak": req.DailyStreak}}
	_, err := m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{Success: true}, nil
}

func (m *UserDataManager) UpdatePlayedGamesCount(req *pb.PlayedGamesCountUReq) (*pb.Void, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": req.UserId}
	update := bson.M{"$set": bson.M{"played_games_count": req.PlayedGamesCount}}
	_, err := m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{Success: true}, nil
}

func (m *UserDataManager) UpdateWinningPercentage(req *pb.WinningPercentageUReq) (*pb.Void, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": req.UserId}
	update := bson.M{"$set": bson.M{"winning_percentage": req.Percentage}}
	_, err := m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{Success: true}, nil
}
