package account_entity

import (
	"time"

	"financo/internal/entities/shared/currencies"
	nullable "financo/internal/entities/shared/nullable_value"
	"github.com/google/uuid"
)

type Entity struct {
	ID       uuid.UUID                 `json:"id"`
	ParentID nullable.Value[uuid.UUID] `json:"parentID"`

	Kind      Kind                `json:"kind"`
	Currency  currencies.Currency `json:"currency"`
	Limit     int64               `json:"limit"`
	IsArchive bool                `json:"isArchive"`

	Name        string                 `json:"name"`
	Description nullable.Value[string] `json:"description"`
	Color       string                 `json:"color"`
	Icon        string                 `json:"icon"`

	CreatedAt time.Time                 `json:"createdAt"`
	UpdatedAt time.Time                 `json:"updatedAt"`
	DeletedAt nullable.Value[time.Time] `json:"-"`

	Children []Entity `json:"-"`
}

func New(
	id uuid.UUID,
	parentID nullable.Value[uuid.UUID],
	kind Kind,
	currency currencies.Currency,
	name string,
	description nullable.Value[string],
	color string,
	icon string,
	limit int64,
	isArchived bool,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt nullable.Value[time.Time],
) Entity {
	var children []Entity

	if parentID.Valid() {
		children = make([]Entity, 0, 10)
	}

	return Entity{
		ID:          id,
		ParentID:    parentID,
		Kind:        kind,
		Currency:    currency,
		Limit:       limit,
		IsArchive:   isArchived,
		Name:        name,
		Description: description,
		Color:       color,
		Icon:        icon,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
		Children:    children,
	}
}
