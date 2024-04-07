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

	defer p.closeDatabaseConnection(conn)

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

// GetChildrenFor finds the children for the given parent Account.
// And returns a copy of the given parent with the children attached to it.
func (p *PostgresRepository) GetChildrenFor(ctx context.Context, parent account.Entity) (account.Entity, error) {
	conn, err := p.service.DB().Conn(ctx)
	if err != nil {
		return parent, errors.Join(ErrInternalServiceFailure, err)
	}

	defer p.closeDatabaseConnection(conn)

	rows, err := conn.QueryContext(
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
			  AND accounts.parent_id = $1
			ORDER BY
			    accounts.is_archived DESC,
			    accounts.created_at ASC
		`,
		parent.ID.String(),
	)
	if err != nil {
		return parent, errors.Join(ErrInternalServiceFailure, err)
	}

	for rows.Next() {
		var (
			record      postgresRecord
			id          uuid.UUID
			description nullable.Value[string]
		)

		err = rows.Scan(
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
		if err != nil {
			return parent, errors.Join(ErrCorruptedAccount, err)
		}

		id, err = uuid.Parse(record.id)
		if err != nil {
			return parent, errors.Join(ErrCorruptedAccount, err)
		}

		if record.description.Valid {
			description, err = postgresParseDescription(record.description)
			if err != nil {
				return parent, errors.Join(ErrCorruptedAccount, err)
			}
		}

		parent.Children = append(
			parent.Children,
			account.New(
				id,
				nullable.New(parent.ID, true),
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
			),
		)
	}

	if err := rows.Err(); err != nil {
		return parent, errors.Join(ErrInternalServiceFailure, err)
	}

	return parent, nil
}

// GetAll finds and returns all not deleted Account records in the system.
// It returns a slice of account_entity.Entity where first it returns all parent
// accounts and later all the child accounts.
func (p *PostgresRepository) GetAll(ctx context.Context) ([]account.Entity, error) {
	accounts := make([]account.Entity, 0, 10)

	conn, err := p.service.DB().Conn(ctx)
	if err != nil {
		return accounts, errors.Join(ErrInternalServiceFailure, err)
	}

	defer p.closeDatabaseConnection(conn)

	rows, err := conn.QueryContext(
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
			ORDER BY
			    accounts.parent_id DESC NULLS FIRST,
			    accounts.is_archived DESC,
			    accounts.created_at ASC
		`,
	)
	if err != nil {
		return accounts, errors.Join(ErrInternalServiceFailure, err)
	}

	for rows.Next() {
		var (
			record      postgresRecord
			id          uuid.UUID
			parentID    nullable.Value[uuid.UUID]
			description nullable.Value[string]
		)

		err = rows.Scan(
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
		if err != nil {
			return accounts, errors.Join(ErrCorruptedAccount, err)
		}

		id, err = uuid.Parse(record.id)
		if err != nil {
			return accounts, errors.Join(ErrCorruptedAccount, err)
		}

		if record.parentID.Valid {
			parentID, err = postgresParseParentID(record.parentID)
			if err != nil {
				return accounts, errors.Join(ErrCorruptedAccount, err)
			}
		}

		if record.description.Valid {
			description, err = postgresParseDescription(record.description)
			if err != nil {
				return accounts, errors.Join(ErrCorruptedAccount, err)
			}
		}

		accounts = append(
			accounts,
			account.New(
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
			),
		)
	}

	if err := rows.Err(); err != nil {
		return accounts, errors.Join(ErrInternalServiceFailure, err)
	}

	return accounts, nil
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

// Delete marks the given account_entity.Entity for deletion in the repository,
// It also marks its children.
func (p *PostgresRepository) Delete(ctx context.Context, id uuid.UUID, time time.Time) error {
	conn, err := p.service.DB().Conn(ctx)
	if err != nil {
		return errors.Join(ErrInternalServiceFailure, err)
	}

	defer p.closeDatabaseConnection(conn)

	result, err := conn.ExecContext(
		ctx,
		`
			UPDATE accounts
			SET accounts.deleted_at = $1
			WHERE id = $2
			   OR parent_id = $2
		`,
		time,
		id.String(),
	)
	if err != nil {
		return errors.Join(ErrInternalServiceFailure, err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Join(ErrInternalServiceFailure, err)
	}

	if rows < 1 {
		return ErrCorruptedAccount
	}

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

// closeDatabaseConnection handles closing the given database connection
func (p *PostgresRepository) closeDatabaseConnection(conn *sql.Conn) {
	if err := conn.Close(); err != nil {
		log.Println("accounts_repository: failed to close postgresql connection:", err)
	}
}
