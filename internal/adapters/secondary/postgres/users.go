package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/NordGus/fnncr/internal/core/domain/user"
	"github.com/NordGus/fnncr/internal/ports"
	"github.com/google/uuid"
)

type UsersRepository struct {
	conn *sql.DB
}

func NewSessionRepository(conn *sql.DB) *UsersRepository {
	return &UsersRepository{
		conn: conn,
	}
}

func (repo *UsersRepository) GetUserByUsername(ctx context.Context, username user.Username) (user.User, error) {
	var (
		uid            string
		passwordDigest string
	)

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}
	defer conn.Close()

	err = conn.
		QueryRowContext(ctx, "SELECT id, password_digest FROM users WHERE username = $1", username.String()).
		Scan(&uid, &passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	id, err := uuid.Parse(uid)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	pwd, err := user.NewPasswordDigest(passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	return user.New(id, username, pwd), nil
}

func (repo *UsersRepository) GetUserByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	var (
		uname          string
		passwordDigest string
	)

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}
	defer conn.Close()

	err = conn.
		QueryRowContext(ctx, "SELECT username, password_digest FROM users WHERE username = $1", id.String()).
		Scan(&uname, &passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	username, err := user.NewUsername(uname)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	pwd, err := user.NewPasswordDigest(passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	return user.New(id, username, pwd), nil
}
