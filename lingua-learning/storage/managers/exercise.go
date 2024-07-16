package managers

import (
	"context"

	pb "learning-service/genprotos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExerciseMongoDBManager struct {
	Collection *mongo.Collection
}

func NewExerciseManager(client *mongo.Client, dbName, collectionName string) *ExerciseMongoDBManager {
	collection := client.Database(dbName).Collection(collectionName)
	return &ExerciseMongoDBManager{Collection: collection}
}

func (m *ExerciseMongoDBManager) Create(exercise *pb.ExerciseCReqUReqForSwagger) (*pb.Void, error) {
	_, err := m.Collection.InsertOne(context.TODO(), exercise)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *ExerciseMongoDBManager) GetByID(req *pb.ByID) (*pb.ExerciseGResUReq, error) {
	var exercise pb.ExerciseGResUReq
	err := m.Collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&exercise)
	if err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (m *ExerciseMongoDBManager) Update(exercise *pb.ExerciseGResUReq) (*pb.Void, error) {
	_, err := m.Collection.UpdateOne(context.TODO(), bson.M{"id": exercise.Id}, bson.D{
		{Key: "$set", Value: exercise},
	})
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *ExerciseMongoDBManager) Delete(req *pb.ByID) (*pb.Void, error) {
	_, err := m.Collection.DeleteOne(context.TODO(), bson.M{"id": req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *ExerciseMongoDBManager) GetAll(req *pb.ExerciseGAReq) (*pb.ExerciseGARes, error) {
	var exercises []*pb.ExerciseGResUReq
	filter := bson.M{}
	if req.LessonId != "" {
		filter["lesson_id"] = req.LessonId
	}
	if req.Type != "" {
		filter["type"] = req.Type
	}

	cursor, err := m.Collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var exercise pb.ExerciseGResUReq
		if err := cursor.Decode(&exercise); err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ExerciseGARes{
		Exercises: exercises,
		Count:     int64(len(exercises)),
	}, nil
}
