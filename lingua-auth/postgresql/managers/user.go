package managers

import (
	"auth-service/models"
	"context"
	"database/sql"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	PgClient    *sql.DB
	MongoClient *mongo.Collection
}

func NewUserManager(db *sql.DB, client *mongo.Client, dbName, collectionName string) *UserManager {
	collection := client.Database(dbName).Collection(collectionName)

	return &UserManager{PgClient: db, MongoClient: collection}
}

func (m *UserManager) Register(req models.RegisterReq) error {
	query := "INSERT INTO users (id, username, email, password, role) VALUES ($1, $2, $3, $4, $5)"
	_, err := m.PgClient.Exec(query, req.ID, req.Username, req.Email, req.Password, req.Role)
	return err
}

func (m *UserManager) RegisterInMongo(id, nativeLang string) error {
	userData := bson.M{
		"userid":            id,
		"nativelang":        nativeLang,
		"level":             "beginner",
		"xp":                0,
		"dailystreak":       0,
		"playedgamescount":  0,
		"winningpercentage": 0.0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.MongoClient.InsertOne(ctx, userData)
	return err
}

func (m *UserManager) Profile(req models.GetProfileReq) (*models.GetProfileResp, error) {
	query := "SELECT id, username, email, password, role FROM users WHERE email = $1"
	row := m.PgClient.QueryRow(query, req.Email)
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
	_, err = m.PgClient.Exec(query, string(hashedPassword), req.Email)
	return err
}

func (m *UserManager) EmailExists(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := m.PgClient.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *UserManager) GetByID(id *models.GetProfileByIdReq) (*models.GetProfileByIdResp, error) {
	query := "SELECT id, username, email, role FROM users WHERE id = $1"
	user := &models.GetProfileByIdResp{}
	err := m.PgClient.QueryRow(query, id.ID).Scan(&user.ID, &user.Username, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserManager) ChangeRole(req models.ChangeRoleReq) error {
	var query string
	if req.ID != "" {
		query = "UPDATE users SET role = $1 WHERE id = $2"
		_, err := m.PgClient.Exec(query, req.Role, req.ID)
		return err
	} else if req.Email != "" {
		query = "UPDATE users SET role = $1 WHERE email = $2"
		_, err := m.PgClient.Exec(query, req.Role, req.Email)
		return err
	}
	return nil
}
