package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"learning-service/config"
	"learning-service/storage/managers"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	DB          *sql.DB
	LessonS     LessonI
	VocabularyS VocabularyI
	ExerciseS   ExerciseI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
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

	// #################     MONGODB CONNECTION     ###################### //
	clientOptions := options.Client().ApplyURI(config.MONGO_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	lm := managers.NewLessonManager(db)
	em := managers.NewExerciseManager(client, config.MONGO_DB_NAME, config.MONGO_COLLECTION_NAME)
	vm := managers.NewVocabularyManager(db)

	log.Println("Successfully connected to the database")
	return &Storage{
		DB:          db,
		LessonS:     lm,
		ExerciseS:   em,
		VocabularyS: vm,
	}, nil
}

// func (s *Storage) Lesson() LessonI {
// 	if s.LessonS == nil {
// 		s.LessonS = managers.NewLessonManager(s.DB)
// 	}
// 	return s.LessonS
// }
