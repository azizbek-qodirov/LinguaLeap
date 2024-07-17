package managers

import (
	"context"
	"database/sql"
	"fmt"

	pb "progress-service/genprotos"
)

type UserLessonManager struct {
	Conn *sql.DB
}

func NewUserLessonManager(conn *sql.DB) *UserLessonManager {
	return &UserLessonManager{Conn: conn}
}

func (m *UserLessonManager) Create(lesson *pb.UserLessonCReq) (*pb.Void, error) {
	query := `INSERT INTO user_lessons (id, user_id, lesson_id) VALUES ($1, $2, $3)`
	_, err := m.Conn.ExecContext(context.Background(), query, lesson.Id, lesson.UserId, lesson.LessonId)
	if err != nil {
		return nil, fmt.Errorf("failed to create user lesson: %w", err)
	}
	return &pb.Void{Success: true}, nil
}
