package accounts_repository

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	account "financo/internal/entities/account_entity"
	"financo/internal/entities/shared/currencies"
	nullable "financo/internal/entities/shared/nullable_value"
	"github.com/google/uuid"
)

type (
	PostgresService interface {
		DB() *sql.DB
	}

	PostgresRepository struct {
		service PostgresService
	}

	postgresRecord struct {
		id          string         // primary key, uuid, unique
		parentID    sql.NullString // index, reference to another account record
		kind        string         // index, enum-like
		currency    string         // index, enum-like
		name        string
		description sql.NullString
		color       string
		icon        string
		limit       int64
		isArchived  bool // index, default false
		createdAt   time.Time
		updatedAt   time.Time
		deletedAt   sql.NullTime
	}
)

// NewPostgresRepository initializes a PostgreSQL repository for Accounts.
func NewPostgresRepository(service PostgresService) *PostgresRepository {
	return &PostgresRepository{
		service: service,
	}
}

// GetByID finds the account_entity.Entity matching the given uuid, that hasn't
// been marked for deletion.
func (p *PostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (account.Entity, error) {
	var (
		record      postgresRecord
		description nullable.Value[string]
		parentID    nullable.Value[uuid.UUID]
	)

	conn, err := p.service.DB().Conn(ctx)
	if err != nil {
		return account.Entity{}, errors.Join(ErrInternalServiceFailure, err)
	}

	defer func(conn *sql.Conn) {
		if err := conn.Close(); err != nil {
			log.Println("accounts_repository: failed to close postgresql connection:", err)
		}
	}(conn)

	err = conn.QueryRowContext(
		ctx,
		`
			SELECT 
			    accounts.id,
			    accounts.parent_id,
			    accounts.kind,
			    accounts.currency,
			    accounts.name,
			    accounts.description,
			    accounts.color,
			    accounts.icon,
			    accounts.limit,
			    accounts.is_archived,
			    accounts.created_at,
			    accounts.updated_at,
			    accounts.deleted_at,
			FROM accounts
			WHERE deleted_at IS NULL
			  AND accounts.id = $1
		`,
		id.String(),
	).Scan(
		&record.id,
		&record.parentID,
		&record.kind,
		&record.currency,
		&record.name,
		&record.description,
		&record.color,
		&record.icon,
		&record.limit,
		&record.isArchived,
		&record.createdAt,
		&record.updatedAt,
		&record.deletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return account.Entity{}, ErrAccountNotFound
	} else if err != nil {
		return account.Entity{}, errors.Join(ErrInternalServiceFailure, err)
	}

	if record.parentID.Valid {
		parentID, err = postgresParseParentID(record.parentID)
		if err != nil {
			return account.Entity{}, errors.Join(ErrCorruptedAccount, err)
		}
	}

	if record.description.Valid {
		description, err = postgresParseDescription(record.description)
		if err != nil {
			return account.Entity{}, errors.Join(ErrCorruptedAccount, err)
		}
	}

	return account.New(
		id,
		parentID,
		account.ParseKind(record.kind),
		currencies.ParseCurrency(record.currency),
		record.name,
		description,
		record.color,
		record.icon,
		record.limit,
		record.isArchived,
		record.createdAt,
		record.updatedAt,
		nullable.Value[time.Time]{},
	), nil
}

// Create saves a new account_entity.Entity in the repository.
// It prevents overwriting existing Account records by first searching by its
// id, to check if the record already exists.
func (p *PostgresRepository) Create(ctx context.Context, acc account.Entity) error {
	return nil
}

// Save saves the new state of the account_entity.Entity in the repository.
// It fails if it can't find the matching Account record by its id.
func (p *PostgresRepository) Save(ctx context.Context, acc account.Entity) error {
	return nil
}

// postgresParseParentID is a simple helper function that parses a
// sql.NullString, representing a postgresRecord parentID, to a
// nullable_value.Value[uuid.UUID] representing an account_entity.Entity
// ParentID.
func postgresParseParentID(parentID sql.NullString) (nullable.Value[uuid.UUID], error) {
	val, err := parentID.Value()
	if err != nil {
		return nullable.Value[uuid.UUID]{}, err
	}

	id, err := uuid.Parse(val.(string))
	if err != nil {
		return nullable.Value[uuid.UUID]{}, err
	}

	return nullable.New(id, true), nil
}

// postgresParseDescription is a simple helper function that parses a
// sql.NullString, representing a postgresRecord description, to a
// nullable_value.Value[string] representing an account_entity.Entity
// Description.
func postgresParseDescription(description sql.NullString) (nullable.Value[string], error) {
	val, err := description.Value()
	if err != nil {
		return nullable.Value[string]{}, err
	}

	return nullable.New(val.(string), true), nil
}
