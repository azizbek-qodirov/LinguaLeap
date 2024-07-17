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

func (m *VocabularyManager) AddTo(req *pb.ByID, exercise *pb.ExerciseGResUReq) (*pb.Void, error) {
	query := `INSERT INTO vocabularies (id, lesson_id, type, question, options, correct_answer) 
	          VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := m.Conn.Exec(query, exercise.Id, exercise.LessonId, exercise.Type, exercise.Question, exercise.Options, exercise.CorrectAnswer)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *VocabularyManager) DeleteFrom(req *pb.ByID) (*pb.Void, error) {
	query := "DELETE FROM vocabularies WHERE id = $1"
	_, err := m.Conn.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *VocabularyManager) Get(req *pb.VocabulariesGAReq) (*pb.VocabulariesGARes, error) {
	query := "SELECT id, lesson_id, type, question, options, correct_answer FROM vocabularies WHERE 1 = 1"
	var args []interface{}
	var paramIndex = 1
	if req.LessonId != "" {
		query += fmt.Sprintf(" AND lesson_id = $%d", paramIndex)
		args = append(args, req.LessonId)
		paramIndex++
	}
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
