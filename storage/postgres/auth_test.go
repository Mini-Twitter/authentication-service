package postgres

import (
	"auth-service/pkg/config"
	"auth-service/pkg/models"
	"fmt"
	"testing"
)

func TestRegister(t *testing.T) {

	configs := config.Load()

	db, err := ConnectPostgres(configs)
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
