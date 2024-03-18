package account_entity

import (
	"time"

	"financo/internal/entities/shared"
	"github.com/google/uuid"
)

type Entity struct {
	ID       uuid.UUID `json:"id"`
	ParentID uuid.UUID `json:"parentID"`

	Type      Type            `json:"type"`
	Currency  shared.Currency `json:"currency"`
	Limit     int64           `json:"limit"`
	IsArchive bool            `json:"isArchive"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Icon        string `json:"icon"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
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
) Entity {
	return Entity{
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
	parent Entity,
	name string,
	description string,
	icon string,
	limit int64,
	isArchived bool,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) Entity {
	return Entity{
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
