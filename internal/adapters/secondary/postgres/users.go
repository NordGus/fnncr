package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/NordGus/fnncr/internal/core/domain/user"
	"github.com/NordGus/fnncr/internal/ports"
	"github.com/google/uuid"
)

var (
	ErrCantConnectToDatabase = errors.New("failed to connect to database")
	ErrCantParseUser         = errors.New("failed to parse user")
)

type usersRepository struct {
	conn *sql.DB
}

func NewUsersRepository(conn *sql.DB) ports.UserRepository {
	return &usersRepository{
		conn: conn,
	}
}

func (repo *usersRepository) GetByUsername(ctx context.Context, username user.Username) (user.User, error) {
	var (
		uid            string
		passwordDigest string
		createdAt      time.Time
		updatedAt      time.Time
	)

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.User{}, errors.Join(ErrCantConnectToDatabase, err)
	}
	defer conn.Close()

	err = conn.
		QueryRowContext(
			ctx,
			"SELECT id, password_digest, created_at, updated_at FROM users WHERE username = $1",
			username.String(),
		).Scan(&uid, &passwordDigest, &createdAt, &updatedAt)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	id, err := uuid.Parse(uid)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	pwd, err := user.NewPasswordDigest(passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	return user.New(id, username, pwd, createdAt, updatedAt), nil
}

func (repo *usersRepository) GetByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	var (
		uname          string
		passwordDigest string
		createdAt      time.Time
		updatedAt      time.Time
	)

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.User{}, errors.Join(ErrCantConnectToDatabase, err)
	}
	defer conn.Close()

	err = conn.
		QueryRowContext(
			ctx,
			"SELECT username, password_digest, created_at, updated_at FROM users WHERE id = $1",
			id.String(),
		).Scan(&uname, &passwordDigest, &createdAt, &updatedAt)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	username, err := user.NewUsername(uname)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	pwd, err := user.NewPasswordDigest(passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	return user.New(id, username, pwd, createdAt, updatedAt), nil
}

func (repo *usersRepository) Create(ctx context.Context, entity user.User) error {
	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return errors.Join(ErrCantConnectToDatabase, err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(
		ctx,
		"INSERT INTO users (id, username, password_digest, created_at, updated_at) VALUES ( $1, $2, $3, $4, $5 )",
		entity.ID.String(),
		entity.Username.String(),
		entity.PasswordDigest.String(),
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	if err != nil {
		return errors.Join(ports.ErrUserNotSaved, err)
	}

	return nil
}
