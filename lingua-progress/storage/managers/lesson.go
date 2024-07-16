package managers

import (
	"database/sql"
	"fmt"
	pb "progress-service/genprotos"
)

type LessonManager struct {
	Conn *sql.DB
}

func NewLessonManager(conn *sql.DB) *LessonManager {
	return &LessonManager{Conn: conn}
}

func (m *LessonManager) Create(lesson *pb.LessonCReqGRes) (*pb.Void, error) {
	query := `INSERT INTO lessons (id, name, title, content, lang_1, lang_2, level, order_number) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := m.Conn.Exec(query, lesson.Id, lesson.Name, lesson.Title, lesson.Content, lesson.Lang_1, lesson.Lang_2, lesson.Level, lesson.Order)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *LessonManager) Update(lesson *pb.LessonUReq) (*pb.Void, error) {
	query := `UPDATE lessons SET name = $1, title = $2, content = $3, level = $4, order_number = $5 WHERE id = $6`
	_, err := m.Conn.Exec(query, lesson.Name, lesson.Title, lesson.Content, lesson.Level, lesson.Order, lesson.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *LessonManager) GetByID(req *pb.ByID) (*pb.LessonCReqGRes, error) {
	query := `SELECT id, name, title, content, lang_1, lang_2, level, order_number FROM lessons WHERE id = $1`
	row := m.Conn.QueryRow(query, req.Id)
	var lesson pb.LessonCReqGRes
	err := row.Scan(&lesson.Id, &lesson.Name, &lesson.Title, &lesson.Content, &lesson.Lang_1, &lesson.Lang_2, &lesson.Level, &lesson.Order)
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (m *LessonManager) Delete(req *pb.ByID) (*pb.Void, error) {
	query := `DELETE FROM lessons WHERE id = $1`
	_, err := m.Conn.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{Success: true}, nil
}

func (m *LessonManager) GetAll(req *pb.LessonGAReq) (*pb.LessonGARes, error) {
	query := `SELECT id, name, title, content, lang_1, lang_2, level, order_number FROM lessons`
	var args []interface{}
	var paramIndex = 1

	conditions := []string{}
	if req.Name != "" {
		conditions = append(conditions, fmt.Sprintf("name = $%d", paramIndex))
		args = append(args, req.Name)
		paramIndex++
	}
	if req.Lang_1 != "" {
		conditions = append(conditions, fmt.Sprintf("lang_1 = $%d", paramIndex))
		args = append(args, req.Lang_1)
		paramIndex++
	}
	if req.Lang_2 != "" {
		conditions = append(conditions, fmt.Sprintf("lang_2 = $%d", paramIndex))
		args = append(args, req.Lang_2)
		paramIndex++
	}
	if req.Level != "" {
		conditions = append(conditions, fmt.Sprintf("level = $%d", paramIndex))
		args = append(args, req.Level)
		paramIndex++
	}
	if req.Order != 0 {
		conditions = append(conditions, fmt.Sprintf("order_number = $%d", paramIndex))
		args = append(args, req.Order)
		paramIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + conditions[0]
		for _, cond := range conditions[1:] {
			query += " AND " + cond
		}
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

	var lessons []*pb.LessonCReqGRes
	for rows.Next() {
		var lesson pb.LessonCReqGRes
		err := rows.Scan(&lesson.Id, &lesson.Name, &lesson.Title, &lesson.Content, &lesson.Lang_1, &lesson.Lang_2, &lesson.Level, &lesson.Order)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, &lesson)
	}

	return &pb.LessonGARes{Lessons: lessons}, nil
}
