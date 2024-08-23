package postgres

import (
	"auth-service/pkg/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"testing"
)

func Connect() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", "5432", "postgres", "123321", "twitter_auth")

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestRegister(t *testing.T) {

	db, err := Connect()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}

	rst := models.RegisterRequest{
		FirstName:   "Hamidulloh",
		LastName:    "Hamidullayev",
		Email:       "hamidulloh@gmail.com",
		Password:    "hamidulloh",
		Phone:       "999747178",
		Username:    "hamidulloh",
		Nationality: "...",
		Bio: "..................." +
			".....................",
	}

	auth := NewAuthRepo(db)

	req, err := auth.Register(rst)
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}

	fmt.Println(req)

}

func TestLoginEmail(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}

	rst := models.LoginEmailRequest{
		Email:    "hamidulloh@gmail.com",
		Password: "hamidulloh",
	}

	auth := NewAuthRepo(db)

	req, err := auth.LoginEmail(rst)
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}

	fmt.Println(req)
}

func TestLoginUsername(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Errorf("Failed to connect to database: %v", err)
	}
	rst := models.LoginUsernameRequest{
		Username: "hamidulloh",
		Password: "hamidulloh",
	}
	auth := NewAuthRepo(db)
	req, err := auth.LoginUsername(rst)
	if err != nil {
		t.Errorf("Failed to register user: %v", err)
	}
	fmt.Println(req)
}
