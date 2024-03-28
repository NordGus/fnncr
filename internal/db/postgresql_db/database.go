package postgresql_db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type (
	Service interface {
		DB() *sql.DB
		Close() error
		Health() map[string]string
	}

	service struct {
		db *sql.DB
	}

	// DBOption is used to inject configuration to the Service
	DBOption func(db *sql.DB)
)

// New returns a new Service. It panics if it can't open a connection to the database.
func New(username string, password string, host string, port string, database string, opts ...DBOption) Service {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	s := &service{db: db}

	for _, applyOption := range opts {
		applyOption(s.db)
	}

	return s
}

func (s *service) DB() *sql.DB {
	return s.db
}

func (s *service) Close() error {
	return s.db.Close()
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		return map[string]string{
			"message": "It's down",
		}
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
