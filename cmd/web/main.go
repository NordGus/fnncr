package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	pgserv "financo/database/postgresql"
	rdsserv "financo/database/redis"
	"financo/internal/adapters/primary/web"
	"financo/internal/adapters/secondary/postgres"
	"financo/internal/adapters/secondary/redis"
	"financo/internal/core/services/authentication"
	_ "github.com/joho/godotenv/autoload"
	goredis "github.com/redis/go-redis/v9"
)

const (
	sessionMaxAge = 7 * 24 * time.Hour
)

func main() {
	var (
		pg = pgserv.New(
			os.Getenv("PG_DB_USERNAME"),
			os.Getenv("PG_DB_PASSWORD"),
			os.Getenv("PG_DB_HOST"),
			os.Getenv("PG_DB_PORT"),
			os.Getenv("PG_DB_DATABASE"),
			func(db *sql.DB) { db.SetMaxOpenConns(10) },
			func(db *sql.DB) { db.SetMaxIdleConns(5) },
			func(db *sql.DB) { db.SetConnMaxIdleTime(15 * time.Second) },
		)
		rds = rdsserv.New(goredis.Options{
			Addr:            "127.0.0.1:6379",
			Password:        "",
			DB:              0,
			PoolSize:        10,
			ConnMaxIdleTime: 1 * time.Second,
		})

		usersrepo   = postgres.NewUsersRepository(pg.DB())
		sessionrepo = redis.NewSessionRepository(rds.Client())

		auth = authentication.NewService(sessionMaxAge, sessionrepo, usersrepo)
	)

	defer pg.Close()
	defer rds.Client()

	app := web.NewApp(
		web.DefaultAppOptions,
		func(a *web.App) { a.AuthAPI = auth },
	)

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
