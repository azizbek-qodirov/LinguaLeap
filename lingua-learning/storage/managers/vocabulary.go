package managers

import (
	"database/sql"
	"fmt"
	pb "learning-service/genprotos"
)

type VocabularyManager struct {
	Conn *sql.DB
}

func NewVocabularyManager(conn *sql.DB) *VocabularyManager {
	return &VocabularyManager{Conn: conn}
}

func (m *VocabularyManager) AddTo(req *pb.ByID) (*pb.Void, error) {
	query := "INSERT INTO vocabulary (exercise_id) VALUES ($1)"
	_, err := m.Conn.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *VocabularyManager) DeleteFrom(req *pb.ByID) (*pb.Void, error) {
	query := "DELETE FROM vocabulary WHERE exercise_id = $1"
	_, err := m.Conn.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *VocabularyManager) Get(req *pb.VocabulariesGAReq) (*pb.VocabulariesGARes, error) {
	query := "SELECT id, lesson_id, type, question, options, correct_answer FROM exercises WHERE lesson_id = $1"
	var args []interface{}
	args = append(args, req.LessonId)
	var paramIndex = 2

	if req.Type != "" {
		query += fmt.Sprintf(" AND type = $%d", paramIndex)
		args = append(args, req.Type)
		paramIndex++
	}

	if req.Pagination != nil {
		if req.Pagination.Limit != 0 {
			query += fmt.Sprintf(" LIMIT $%d", paramIndex)
			args = append(args, req.Pagination.Limit)
			paramIndex++
		}
		if req.Pagination.Offset != 0 {
			query += fmt.Sprintf(" OFFSET $%d", paramIndex)
			args = append(args, req.Pagination.Offset)
			paramIndex++
		}
	}

	rows, err := m.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vocabularies []*pb.ExerciseGResUReq
	for rows.Next() {
		var vocabulary pb.ExerciseGResUReq
		err := rows.Scan(&vocabulary.Id, &vocabulary.LessonId, &vocabulary.Type, &vocabulary.Question, &vocabulary.Options, &vocabulary.CorrectAnswer)
		if err != nil {
			return nil, err
		}
		vocabularies = append(vocabularies, &vocabulary)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.VocabulariesGARes{Vocabularies: vocabularies}, nil
}
