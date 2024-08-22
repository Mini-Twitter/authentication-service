package postgres

import (
	config "auth-service/pkg/config"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func ConnectPostgres(config config.Config) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DB_HOST, config.DB_PORT, config.USER_PORT, config.DB_PASSWORD, config.DB_NAME)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
