package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/NordGus/fnncr/database/postgresql"
	rdsdb "github.com/NordGus/fnncr/database/redis"
	"github.com/NordGus/fnncr/internal/adapters/primary/web"
	"github.com/NordGus/fnncr/internal/adapters/secundary/postgres"
	rdsadapter "github.com/NordGus/fnncr/internal/adapters/secundary/redis"
	"github.com/NordGus/fnncr/internal/core/services/authentication"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

func main() {
	var (
		pg = postgresql.New(
			"fnncr", "local_dev", "127.0.0.1", 5432, "fnncr_dev",
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

		usersrepo   = postgres.NewSessionRepository(pg.DB())
		sessionrepo = rdsadapter.NewSessionRepository(rds.Client())

		auth = authentication.NewService(sessionrepo, usersrepo)
	)

	defer pg.Close()
	defer rds.Client()

	app := web.NewApp(
		web.DefaultAppOptions,
		func(a *web.App) { a.AuthAPI = auth },
	)

	app.SetRoutes()

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
