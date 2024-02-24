package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	web "financo/internal/adapters/primary/web/api"
	pgadpter "financo/internal/adapters/secondary/postgres"
	rdsadapter "financo/internal/adapters/secondary/redis"
	"financo/internal/core/services/authentication"
	pgdb "financo/internal/database/postgresql"
	rdsdb "financo/internal/database/redis"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

const (
	sessionMaxAge = 7 * 24 * time.Hour
)

func main() {
	var (
		pg = pgdb.New(
			os.Getenv("PG_DB_USERNAME"),
			os.Getenv("PG_DB_PASSWORD"),
			os.Getenv("PG_DB_HOST"),
			os.Getenv("PG_DB_PORT"),
			os.Getenv("PG_DB_DATABASE"),
			func(db *sql.DB) { db.SetMaxOpenConns(10) },
			func(db *sql.DB) { db.SetMaxIdleConns(5) },
			func(db *sql.DB) { db.SetConnMaxIdleTime(15 * time.Second) },
		)
		rds = rdsdb.New(redis.Options{
			Addr:            "127.0.0.1:6379",
			Password:        "",
			DB:              0,
			PoolSize:        10,
			ConnMaxIdleTime: 1 * time.Second,
		})

		usersRepo   = pgadpter.NewUsersRepository(pg.DB())
		sessionRepo = rdsadapter.NewSessionRepository(rds.Client())

		authAPI = authentication.NewService(sessionMaxAge, sessionRepo, usersRepo)
	)

	defer pg.Close()
	defer rds.Client()

	app := web.NewApp(
		web.DefaultAppOptions,
		func(a *web.App) { a.AuthAPI = authAPI },
	)

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
