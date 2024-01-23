package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	pgserv "github.com/NordGus/fnncr/database/postgresql"
	rdsserv "github.com/NordGus/fnncr/database/redis"
	"github.com/NordGus/fnncr/internal/adapters/primary/web"
	"github.com/NordGus/fnncr/internal/adapters/secondary/postgres"
	"github.com/NordGus/fnncr/internal/adapters/secondary/redis"
	"github.com/NordGus/fnncr/internal/core/services/authentication"
	_ "github.com/joho/godotenv/autoload"
	goredis "github.com/redis/go-redis/v9"
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

		auth = authentication.NewService(sessionrepo, usersrepo)
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
