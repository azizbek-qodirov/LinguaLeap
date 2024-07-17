package managers

import (
	"database/sql"
	pb "progress-service/genprotos"
)

type UserLessonManager struct {
	Conn *sql.DB
}

func NewUserLessonManager(conn *sql.DB) *UserLessonManager {
	return &UserLessonManager{Conn: conn}
}

func (m *UserLessonManager) Create(lesson *pb.UserLessonCReq) (*pb.Void, error) {
	return nil, nil
}
