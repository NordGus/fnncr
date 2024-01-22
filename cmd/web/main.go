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
		assets = os.DirFS("./dist")

		pg = pgserv.New(
			"fnncr", "local_dev", "127.0.0.1", 5432, "fnncr_dev",
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

		usersrepo   = postgres.NewSessionRepository(pg.DB())
		sessionrepo = redis.NewSessionRepository(rds.Client())

		auth = authentication.NewService(sessionrepo, usersrepo)
	)

	defer pg.Close()
	defer rds.Client()

	app := web.NewApp(
		web.DefaultAppOptions,
		func(a *web.App) { a.AssetsFS = assets },
		func(a *web.App) { a.AuthAPI = auth },
	)

	app.SetRoutes()

	if err := app.Run(); err != nil {
		log.Fatalln(err)
	}
}
