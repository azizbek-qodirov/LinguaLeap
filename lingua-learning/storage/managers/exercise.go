package managers

import (
	"context"
	"fmt"

	pb "learning-service/genprotos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type ExerciseTemp struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	LessonId      string             `bson:"lessonid"`
	Type          string             `bson:"type"`
	Question      string             `bson:"question"`
	Options       string             `bson:"options"`
	CorrectAnswer string             `bson:"correctanswer"`
}

type ExerciseMongoDBManager struct {
	Collection *mongo.Collection
}

func NewExerciseManager(client *mongo.Client, dbName, collectionName string) *ExerciseMongoDBManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &ExerciseMongoDBManager{Collection: collection}
}

func (m *ExerciseMongoDBManager) Create(exercise *pb.ExerciseCReqUReqForSwagger) (*pb.Void, error) {
	_, err := m.Collection.InsertOne(ctx, exercise)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *ExerciseMongoDBManager) GetByID(req *pb.ByID) (*pb.ExerciseGResUReq, error) {
	var temp ExerciseTemp
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}
	err = m.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&temp)
	if err != nil {
		return nil, err
	}
	exercise := &pb.ExerciseGResUReq{
		Id:            temp.Id.Hex(),
		LessonId:      temp.LessonId,
		Type:          temp.Type,
		Question:      temp.Question,
		Options:       temp.Options,
		CorrectAnswer: temp.CorrectAnswer,
	}
	return exercise, nil
}

func (m *ExerciseMongoDBManager) Update(exercise *pb.ExerciseGResUReq) (*pb.Void, error) {
	id, err := primitive.ObjectIDFromHex(exercise.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}
	_, err = m.Collection.UpdateOne(ctx, bson.M{"_id": id}, bson.D{
		{Key: "$set", Value: exercise},
	})
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *ExerciseMongoDBManager) Delete(req *pb.ByID) (*pb.Void, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}
	result, err := m.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("no document found with the given id")
	}
	return &pb.Void{Success: true}, nil
}

func (m *ExerciseMongoDBManager) GetAll(req *pb.ExerciseGAReq) (*pb.ExerciseGARes, error) {
	var exercises []*pb.ExerciseGResUReq
	filter := bson.M{}
	if req.LessonId != "" {
		filter["lessonid"] = req.LessonId
	}
	if req.Type != "" {
		filter["type"] = req.Type
	}

	findOptions := options.Find()
	if req.Pagination != nil {
		if req.Pagination.Limit > 0 {
			findOptions.SetLimit(req.Pagination.Limit)
		}
		if req.Pagination.Offset > 0 {
			findOptions.SetSkip(req.Pagination.Offset)
		}
	}

	cursor, err := m.Collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var temp ExerciseTemp
		if err := cursor.Decode(&temp); err != nil {
			return nil, err
		}
		exercise := &pb.ExerciseGResUReq{
			Id:            temp.Id.Hex(),
			LessonId:      temp.LessonId,
			Type:          temp.Type,
			Question:      temp.Question,
			Options:       temp.Options,
			CorrectAnswer: temp.CorrectAnswer,
		}
		exercises = append(exercises, exercise)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	totalCount, err := m.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &pb.ExerciseGARes{
		Exercises: exercises,
		Count:     totalCount,
	}, nil
}
