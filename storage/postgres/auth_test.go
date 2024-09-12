package postgres

import (
	"auth-service/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

// Connect establishes a connection to the PostgreSQL database.
func Connect() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "123321", "twitter_auth")

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// TestRegister tests the Register method of AuthRepo.
func TestRegister(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	rst := models.RegisterRequest{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		Password:    "password123",
		Phone:       "1234567890",
		Username:    "johndoe",
		Nationality: "American",
		Bio:         "Software Developer",
	}

	auth := NewAuthRepo(db)

	req, err := auth.Register(rst)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	if req.Email != rst.Email {
		t.Errorf("Expected email %v, got %v", rst.Email, req.Email)
	}

	fmt.Println("Register response:", req)
}

// TestLoginEmail tests the LoginEmail method of AuthRepo.
func TestLoginEmail(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	rst := models.LoginEmailRequest{
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	auth := NewAuthRepo(db)

	req, err := auth.LoginEmail(rst)
	if err != nil {
		t.Fatalf("Failed to login using email: %v", err)
	}

	if req.Email != rst.Email {
		t.Errorf("Expected email %v, got %v", rst.Email, req.Email)
	}

	fmt.Println("LoginEmail response:", req)
}

//// TestLoginUsername tests the LoginUsername method of AuthRepo.
//func TestLoginUsername(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	rst := models.LoginUsernameRequest{
//		Username: "johndoe",
//		Password: "password123",
//	}
//
//	auth := NewAuthRepo(db)
//
//	req, err := auth.LoginUsername(rst)
//	if err != nil {
//		t.Fatalf("Failed to login using username: %v", err)
//	}
//
//	if req.Username != rst.Username {
//		t.Errorf("Expected username %v, got %v", rst.Username, req.Username)
//	}
//
//	fmt.Println("LoginUsername response:", req)
//}
