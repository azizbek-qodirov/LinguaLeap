package storage

import (
	"context"
	"database/sql"
	"fmt"

	"progress-service/config"
	"progress-service/storage/managers"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Storage struct {
	PgClient            *sql.DB
	MongoClient         *mongo.Client
	LessonServiceClient *grpc.ClientConn

	UserLessonS UserLessonI
	UserDataS   UserDataI
	QuizS       QuizI
}

func NewPostgresStorage(config config.Config, lessonServiceClient *grpc.ClientConn) (*Storage, error) {
	// #################    POSTGRESQL CONNECTION     ###################### //
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to the database pgsql!!!")

	// #################     MONGODB CONNECTION     ###################### //
	clientOptions := options.Client().ApplyURI(config.MONGO_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to the database mongodb!!!")

	ul := managers.NewUserLessonManager(db)
	ud := managers.NewUserDataManager(client, config.MONGO_DB_NAME, config.MONGO_COLLECTION_NAME, db)

	return &Storage{
		PgClient:            db,
		UserLessonS:         ul,
		LessonServiceClient: lessonServiceClient,
		UserDataS:           ud,
	}, nil
}

func (s *Storage) UserData() UserDataI {
	if s.UserDataS == nil {
		s.UserDataS = managers.NewUserDataManager(s.MongoClient, config.Load().MONGO_DB_NAME, config.Load().MONGO_COLLECTION_NAME, s.PgClient)
	}
	return s.UserDataS
}

func (s *Storage) UserLesson() UserLessonI {
	if s.UserLessonS == nil {
		s.UserLessonS = managers.NewUserLessonManager(s.PgClient)
	}
	return s.UserLessonS
}

func (s *Storage) Quiz() QuizI {
	if s.QuizS == nil {
		s.QuizS = managers.NewQuizManager(s.PgClient)
	}
	return s.QuizS
}
