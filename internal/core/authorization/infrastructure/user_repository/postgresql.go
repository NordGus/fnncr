package user_repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"financo/internal/core/authorization"
	"financo/internal/core/authorization/domain/passworddigest"
	"financo/internal/core/authorization/domain/sessionversion"
	"financo/internal/core/authorization/domain/timestamp"
	"financo/internal/core/authorization/domain/user"
	"financo/internal/core/authorization/domain/userID"
	"financo/internal/core/authorization/domain/username"
)

type (
	PostgresService interface {
		DB() *sql.DB
	}

	postgreSQLRepository struct {
		conn                *sql.DB
		userIDEncoder       userID.Encoder
		passwordDigestCrypt passworddigest.Crypt
	}

	postgreSQLRecord struct {
		id             string
		username       string
		passwordDigest string
		sessionVersion uint32
		createdAt      time.Time
		updatedAt      time.Time
	}
)

func NewPostgreSQLRepository(
	serv PostgresService,
	userIDEncoder userID.Encoder,
	pwdCrypt passworddigest.Crypt,
) authorization.UserRepository {
	return &postgreSQLRepository{
		conn:                serv.DB(),
		userIDEncoder:       userIDEncoder,
		passwordDigestCrypt: pwdCrypt,
	}
}

func (repo *postgreSQLRepository) GetByID(ctx context.Context, value userID.Value) (user.Entity, error) {
	var record postgreSQLRecord

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserNotFound, err)
	}
	defer conn.Close()

	err = conn.QueryRowContext(
		ctx,
		"SELECT users.id, users.username, users.password_digest, users.session_version, users.created_at, users.updated_at FROM users WHERE users.id = $1 AND users.deleted_at IS NULL",
		value.String(),
	).Scan(
		&record.id,
		&record.username,
		&record.passwordDigest,
		&record.sessionVersion,
		&record.createdAt,
		&record.updatedAt,
	)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserNotFound, err)
	}

	uname, err := username.New(record.username)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	pwd, err := passworddigest.New(record.passwordDigest, repo.passwordDigestCrypt)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	ver, err := sessionversion.New(record.sessionVersion)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	ct, err := timestamp.New(record.createdAt)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	ut, err := timestamp.New(record.updatedAt)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	return user.New(value, uname, pwd, ver, ct, ut), nil
}

func (repo *postgreSQLRepository) GetByUsername(ctx context.Context, value username.Value) (user.Entity, error) {
	var record postgreSQLRecord

	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserNotFound, err)
	}
	defer conn.Close()

	err = conn.QueryRowContext(
		ctx,
		"SELECT users.id, users.username, users.password_digest, users.session_version, users.created_at, users.updated_at FROM users WHERE users.username = $1 AND users.deleted_at IS NULL",
		value.String(),
	).Scan(
		&record.id,
		&record.username,
		&record.passwordDigest,
		&record.sessionVersion,
		&record.createdAt,
		&record.updatedAt,
	)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserNotFound, err)
	}

	id, err := userID.New(record.id, repo.userIDEncoder)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	pwd, err := passworddigest.New(record.passwordDigest, repo.passwordDigestCrypt)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	ver, err := sessionversion.New(record.sessionVersion)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	ct, err := timestamp.New(record.createdAt)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	ut, err := timestamp.New(record.updatedAt)
	if err != nil {
		return user.Entity{}, errors.Join(ErrUserCantBeParsed, err)
	}

	return user.New(id, value, pwd, ver, ct, ut), nil
}

func (repo *postgreSQLRepository) Save(ctx context.Context, entity user.Entity) error {
	conn, err := repo.conn.Conn(ctx)
	if err != nil {
		return errors.Join(ErrUserWasNotSaved, err)
	}
	defer conn.Close()

	result, err := conn.ExecContext(
		ctx,
		"UPDATE users SET session_version = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL",
		entity.SessionVersion().Uint32(),
		entity.UpdatedAt().Time(),
		entity.ID().String(),
	)
	if err != nil {
		return errors.Join(ErrUserWasNotSaved, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Join(ErrUserWasNotSaved, err)
	}

	if rows < 1 {
		return ErrUserWasNotSaved
	}

	if rows > 1 {
		return ErrUserRepositoryIntegrityCorrupted
	}

	return nil
}
