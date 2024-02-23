package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"financo/internal/core/domain/user"
	"financo/internal/ports"
	"github.com/google/uuid"
)

var (
	ErrCantConnectToDatabase = errors.New("users repository: failed to connect to database")
	ErrCantParseUser         = errors.New("users repository: failed to parse user")
)

type (
	usersRepository struct {
		conn *sql.DB
	}

	record struct {
		id             string
		username       string
		passwordDigest string
		sessionVersion int32
		createdAt      time.Time
		updatedAt      time.Time
		deletedAt      sql.NullTime
	}
)

func NewUsersRepository(conn *sql.DB) ports.UserRepository {
	return &usersRepository{
		conn: conn,
	}
}

func (repo *usersRepository) GetByUsername(ctx context.Context, username user.Username) (user.User, error) {
	var rcrd record

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.User{}, errors.Join(ErrCantConnectToDatabase, err)
	}
	defer conn.Close()

	err = conn.
		QueryRowContext(
			ctx,
			"SELECT id, username, password_digest, session_version, created_at, updated_at, deleted_at FROM users WHERE username = $1 AND deleted_at IS NULL",
			username.String(),
		).
		Scan(
			&rcrd.id,
			&rcrd.username,
			&rcrd.passwordDigest,
			&rcrd.sessionVersion,
			&rcrd.createdAt,
			&rcrd.updatedAt,
			&rcrd.deletedAt,
		)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	id, err := uuid.Parse(rcrd.id)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	pwd, err := user.NewPasswordDigest(rcrd.passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	return user.New(id, username, pwd, rcrd.sessionVersion, rcrd.createdAt, rcrd.updatedAt), nil
}

func (repo *usersRepository) GetByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	var rcrd record

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.User{}, errors.Join(ErrCantConnectToDatabase, err)
	}
	defer conn.Close()

	err = conn.
		QueryRowContext(
			ctx,
			"SELECT id, username, password_digest, session_version, created_at, updated_at, deleted_at FROM users WHERE id = $1 AND deleted_at IS NULL",
			id.String(),
		).
		Scan(
			&rcrd.id,
			&rcrd.username,
			&rcrd.passwordDigest,
			&rcrd.sessionVersion,
			&rcrd.createdAt,
			&rcrd.updatedAt,
			&rcrd.deletedAt,
		)
	if err != nil {
		return user.User{}, errors.Join(ports.ErrUserNotFound, err)
	}

	username, err := user.NewUsername(rcrd.username)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	pwd, err := user.NewPasswordDigest(rcrd.passwordDigest)
	if err != nil {
		return user.User{}, errors.Join(ErrCantParseUser, err)
	}

	return user.New(id, username, pwd, rcrd.sessionVersion, rcrd.createdAt, rcrd.updatedAt), nil
}

func (repo *usersRepository) Create(ctx context.Context, entity user.User) error {
	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return errors.Join(ErrCantConnectToDatabase, err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(
		ctx,
		"INSERT INTO users (id, username, password_digest, session_version, created_at, updated_at) VALUES ( $1, $2, $3, $4, $5, $6 )",
		entity.ID.String(),
		entity.Username.String(),
		entity.PasswordDigest.String(),
		entity.SessionVersion,
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	if err != nil {
		return errors.Join(ports.ErrUserNotSaved, err)
	}

	return nil
}

func (repo *usersRepository) Update(ctx context.Context, entity user.User) (user.User, error) {
	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return entity, errors.Join(ErrCantConnectToDatabase, err)
	}
	defer conn.Close()

	_, err = conn.ExecContext(
		ctx,
		"UPDATE users SET id = $1, username = $2, password_digest = $3, session_version = $4, created_at = $5, updated_at = $6 WHERE id = $7",
		entity.ID.String(),
		entity.Username.String(),
		entity.PasswordDigest.String(),
		entity.SessionVersion,
		entity.CreatedAt,
		entity.UpdatedAt,
		entity.ID.String(),
	)
	if err != nil {
		return entity, errors.Join(ports.ErrUserNotSaved, err)
	}

	return entity, nil
}
