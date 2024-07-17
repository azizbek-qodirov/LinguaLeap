package managers

import (
	"auth-service/models"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB
var userManager *UserManager
var user_id string

func TestMain(m *testing.M) {
	fmt.Println("Connecting to the database...")
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable", "localhost", "mrbek", "auth_db", "QodirovCoder", 5432)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("could not connect to the DB: %v", err)
	}

	userManager = NewUserManager(db)

	fmt.Println("Database connected!!!")
	code := m.Run()

	fmt.Println("Database connection closed")
	db.Close()
	os.Exit(code)
}

func TestRegister(t *testing.T) {
	fmt.Println("Testing user registration...")
	user_id = uuid.NewString()
	req := models.RegisterReq{
		ID:       user_id,
		Username: "testuser",
		Email:    "testuser@example.com",
		Password: "testpassword",
		Role:     "user",
	}

	err := userManager.Register(req)
	assert.NoError(t, err)

	exists, err := userManager.EmailExists(req.Email)
	assert.NoError(t, err)
	assert.True(t, exists)
	fmt.Println("OK. User registered successfully")
}

func TestProfile(t *testing.T) {
	fmt.Println("Testing user profile retrieval...")
	req := models.GetProfileReq{
		Email: "testuser@example.com",
	}

	profile, err := userManager.Profile(req)
	assert.NoError(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, req.Email, profile.Email)
	fmt.Println("OK. User profile retrieved successfully")
}

func TestEmailExists(t *testing.T) {
	fmt.Println("Testing email existence check...")
	email := "testuser@example.com"

	exists, err := userManager.EmailExists(email)
	assert.NoError(t, err)
	assert.True(t, exists)
	fmt.Println("OK. Email existence check successful")
}

func TestGetByID(t *testing.T) {
	fmt.Println("Testing get user by ID...")
	getByIDReq := models.GetProfileByIdReq{
		ID: user_id,
	}

	profile, err := userManager.GetByID(&getByIDReq)
	assert.NoError(t, err)
	assert.NotNil(t, profile)
	assert.Equal(t, user_id, profile.ID)
	fmt.Println("OK. Get user by ID successful")
}

func TestChangeRole(t *testing.T) {
	fmt.Println("Testing change user role...")
	req := models.ChangeRoleReq{
		Email: "testuser@example.com",
		Role:  "admin",
	}

	err := userManager.ChangeRole(req)
	assert.NoError(t, err)

	// Verify role change
	profileReq := models.GetProfileReq{
		Email: req.Email,
	}
	profile, err := userManager.Profile(profileReq)
	assert.NoError(t, err)
	assert.Equal(t, req.Role, profile.Role)
	fmt.Println("OK. User role changed successfully")
}
