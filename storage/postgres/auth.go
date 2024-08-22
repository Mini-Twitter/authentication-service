package postgres

import (
	"auth-service/pkg/models"
	"auth-service/storage"
	"github.com/jmoiron/sqlx"
	"time"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) storage.AuthStorage {
	return &AuthRepo{
		db: db,
	}
}

func (a *AuthRepo) Register(in models.RegisterRequest) (models.RegisterResponse, error) {

	var id string
	query := `INSERT INTO users (phone, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := a.db.QueryRow(query, in.Phone, in.Email, in.Password, time.Now().String()).Scan(&id)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	query1 := `INSERT INTO user_profile (user_id, first_name, last_name, username, nationality, bio, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = a.db.Query(query1, id, in.FirstName, in.LastName, in.Username, in.Nationality, in.Bio, time.Now().String())
	if err != nil {
		return models.RegisterResponse{}, err
	}

	return models.RegisterResponse{
		Id:          id,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Email:       in.Email,
		Phone:       in.Phone,
		Username:    in.Username,
		Nationality: in.Nationality,
		Bio:         in.Bio,
	}, nil
}
func (a *AuthRepo) LoginEmail(in models.LoginEmailRequest) (models.LoginEmailResponse, error) {
	res := models.LoginEmailResponse{}

	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := a.db.Get(&res, query, in.Email)
	if err != nil {
		return models.LoginEmailResponse{}, err
	}

	var role string
	query1 := `SELECT role FROM user_profile WHERE user_id = $1`
	err = a.db.Get(&role, query1, res.Id)

	return models.LoginEmailResponse{
		Id:       res.Id,
		Email:    res.Email,
		Role:     role,
		Password: res.Password,
	}, nil
}
func (a *AuthRepo) LoginUsername(in models.LoginUsernameRequest) (models.LoginUsernameResponse, error) {

	var id string
	var role string
	query := `SELECT role, user_id FROM user_profile WHERE username = $1`
	err := a.db.QueryRow(query, in.Username).Scan(&role, &id)
	if err != nil {
		return models.LoginUsernameResponse{}, err
	}

	var password string
	query1 := `SELECT password FROM users WHERE id = $1`
	err = a.db.Get(&password, query1, id)
	if err != nil {
		return models.LoginUsernameResponse{}, err
	}

	return models.LoginUsernameResponse{
		Id:       id,
		Username: in.Username,
		Password: password,
		Role:     role,
	}, nil
}
