package account_model

import (
	"financo/internal/http/rest/models/shared"
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID       uuid.UUID
	ParentID uuid.UUID

	Type      Type
	Currency  shared.Currency
	Limit     int64
	IsArchive bool

	Name        string
	Description string
	Color       string
	Icon        string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func New(
	id uuid.UUID,
	t Type,
	c shared.Currency,
	name string,
	description string,
	color string,
	icon string,
	limit int64,
	isArchived bool,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) Model {
	return Model{
		ID:          id,
		ParentID:    uuid.UUID{},
		Type:        t,
		Currency:    c,
		Limit:       limit,
		IsArchive:   isArchived,
		Name:        name,
		Description: description,
		Color:       color,
		Icon:        icon,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
	}
}

func NewChild(
	id uuid.UUID,
	parent Model,
	name string,
	description string,
	icon string,
	limit int64,
	isArchived bool,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) Model {
	return Model{
		ID:          id,
		ParentID:    parent.ID,
		Type:        parent.Type,
		Currency:    parent.Currency,
		Limit:       limit,
		IsArchive:   isArchived,
		Name:        name,
		Description: description,
		Color:       parent.Color,
		Icon:        icon,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
	}
}
