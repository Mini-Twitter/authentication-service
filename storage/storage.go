package storage

import "auth-service/pkg/models"

type AuthStorage interface {
	Register(in models.RegisterRequest) (models.RegisterResponse, error)
	LoginEmail(in models.LoginEmailRequest) (models.LoginEmailResponse, error)
	LoginUsername(in models.LoginUsernameRequest) (models.LoginUsernameResponse, error)
}
