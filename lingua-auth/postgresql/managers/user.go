package managers

import (
	"auth-service/models"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	Conn *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{Conn: db}
}

func (m *UserManager) Register(req models.RegisterReq) error {
	query := "INSERT INTO users (id, username, email, password, role) VALUES ($1, $2, $3, $4, $5)"
	_, err := m.Conn.Exec(query, req.ID, req.Username, req.Email, req.Password, req.Role)
	return err
}

func (m *UserManager) Profile(req models.GetProfileReq) (*models.GetProfileResp, error) {
	query := "SELECT id, username, email, password, role FROM users WHERE email = $1"
	row := m.Conn.QueryRow(query, req.Email)
	var user models.GetProfileResp
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserManager) UpdatePassword(req *models.UpdatePasswordReq) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	query := "UPDATE users SET password = $1 WHERE email = $2"
	_, err = m.Conn.Exec(query, string(hashedPassword), req.Email)
	return err
}

func (m *UserManager) EmailExists(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := m.Conn.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *UserManager) GetByID(id *models.GetProfileByIdReq) (*models.GetProfileByIdResp, error) {
	query := "SELECT id, username, email, role FROM users WHERE id = $1"
	user := &models.GetProfileByIdResp{}
	err := m.Conn.QueryRow(query, id.ID).Scan(&user.ID, &user.Username, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserManager) ChangeRole(req models.ChangeRoleReq) error {
	var query string
	if req.ID != "" {
		query = "UPDATE users SET role = $1 WHERE id = $2"
		_, err := m.Conn.Exec(query, req.Role, req.ID)
		return err
	} else if req.Email != "" {
		query = "UPDATE users SET role = $1 WHERE email = $2"
		_, err := m.Conn.Exec(query, req.Role, req.Email)
		return err
	}
	return nil
}
