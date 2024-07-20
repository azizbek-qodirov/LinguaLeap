package managers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "progress-service/genprotos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDataManager struct {
	Collection *mongo.Collection
	PgClient   *sql.DB
}

func NewUserDataManager(client *mongo.Client, dbName, collectionName string, pgClient *sql.DB) *UserDataManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &UserDataManager{Collection: collection, PgClient: pgClient}
}

func (m *UserDataManager) GetUserData(req *pb.ByID) (*pb.UserDataGRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result pb.UserDataGRes
	filter := bson.M{"userid": req.Id}
	err := m.Collection.FindOne(ctx, filter).Decode(&result)
	fmt.Println(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *UserDataManager) UpdateXP(req *pb.XPUReq) (*pb.Void, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"userid": req.UserId}
	update := bson.M{"$inc": bson.M{"xp": req.Xp}}
	_, err := m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{Success: true}, nil
}

func (m *UserDataManager) UpdateDailyStreak(req *pb.StreakUReq) (*pb.Void, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check PostgreSQL for lessons today
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	var count int
	err := m.PgClient.QueryRow(`SELECT COUNT(*) FROM user_lessons WHERE user_id = $1 AND DATE(created_at) = $2`, req.UserId, today).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count > 0 {
		// User has lessons today, do nothing
		return &pb.Void{Success: true}, nil
	}

	// Check PostgreSQL for lessons yesterday
	err = m.PgClient.QueryRow(`SELECT COUNT(*) FROM user_lessons WHERE user_id = $1 AND DATE(created_at) = $2`, req.UserId, yesterday).Scan(&count)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"userid": req.UserId}
	if count == 0 {
		// No lessons yesterday, reset streak to 0
		update := bson.M{"$set": bson.M{"dailystreak": 0}}
		_, err := m.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return nil, err
		}
	} else {
		// Lessons yesterday, increment streak by 1
		update := bson.M{"$inc": bson.M{"dailystreak": 1}}
		_, err := m.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Void{Success: true}, nil
}

func (m *UserDataManager) UpdateWinningPercentage(req *pb.WinningPercentageUReq) (*pb.Void, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Retrieve the current winning percentage and played games count
	filter := bson.M{"userid": req.UserId}
	var userData pb.UserDataGRes
	err := m.Collection.FindOne(ctx, filter).Decode(&userData)
	if err != nil {
		return nil, err
	}

	// Calculate the new winning percentage
	newWinningPercentage := (userData.WinningPercentage*float32(userData.PlayedGamesCount) + req.Percentage) / float32(userData.PlayedGamesCount+1)

	// Update the winning percentage and increment the played games count
	update := bson.M{
		"$set": bson.M{"winningpercentage": newWinningPercentage},
		"$inc": bson.M{"playedgamescount": 1},
	}
	_, err = m.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.Void{Success: true}, nil
}

func (m *UserDataManager) GetLeadBoard(req *pb.Void) (*pb.LeadboardRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define an empty filter to get all documents
	filter := bson.M{}

	// Define options to sort by XP in descending order
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "xp", Value: -1}})

	// Perform the find operation
	cursor, err := m.Collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Initialize a slice to hold the user data
	var users []*pb.LeadboardUserRes

	// Iterate over the cursor to decode each document into the UserDataGRes struct
	for cursor.Next(ctx) {
		var user pb.LeadboardUserRes
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Create the leaderboard response
	leadboardRes := &pb.LeadboardRes{
		Users: users,
	}

	return leadboardRes, nil
}
