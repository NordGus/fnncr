package postgres

import (
	"context"
	"database/sql"
	"errors"

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
	)

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.User{}, errors.Join(ErrCantConnectToDatabase, err)
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
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	pwd, err := user.NewPasswordDigest(passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	return user.New(id, username, pwd), nil
}

func (repo *usersRepository) GetByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	var (
		uname          string
		passwordDigest string
	)

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.User{}, errors.Join(ErrCantConnectToDatabase, err)
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
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	pwd, err := user.NewPasswordDigest(passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	return user.New(id, username, pwd), nil
}
