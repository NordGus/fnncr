package account_form

import (
	"time"

	account "financo/internal/entities/account_entity"
	"financo/internal/entities/shared/currencies"
	"github.com/google/uuid"
)

type Form struct {
	ParentID ParentID `json:"parentID"`

	Kind      Kind      `json:"kind"`
	Currency  Currency  `json:"currency"`
	Limit     Limit     `json:"limit"`
	IsArchive IsArchive `json:"isArchive"`

	Name        Name        `json:"name"`
	Description Description `json:"description"`
	Color       Color       `json:"color"`
	Icon        Icon        `json:"icon"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}

func New(
	t account.Kind,
	c currencies.Currency,
	name string,
	description string,
	color string,
	icon string,
	limit int64,
	isArchived bool,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) Form {
	f := Form{
		Kind:        Kind{Value: t, Errors: nil},
		Currency:    Currency{Value: c, Errors: nil},
		Limit:       Limit{Value: limit, Errors: nil},
		IsArchive:   IsArchive{Value: isArchived, Errors: nil},
		Name:        Name{Value: name, Errors: nil},
		Description: Description{Value: description, Errors: nil},
		Color:       Color{Value: color, Errors: nil},
		Icon:        Icon{Value: icon, Errors: nil},
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
	}

	return f
}

func NewChild(
	parentID uuid.UUID,
	t account.Kind,
	c currencies.Currency,
	name string,
	description string,
	color string,
	icon string,
	limit int64,
	isArchived bool,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
) Form {
	f := Form{
		ParentID:    ParentID{Value: parentID, Errors: nil},
		Kind:        Kind{Value: t, Errors: nil},
		Currency:    Currency{Value: c, Errors: nil},
		Limit:       Limit{Value: limit, Errors: nil},
		IsArchive:   IsArchive{Value: isArchived, Errors: nil},
		Name:        Name{Value: name, Errors: nil},
		Description: Description{Value: description, Errors: nil},
		Color:       Color{Value: color, Errors: nil},
		Icon:        Icon{Value: icon, Errors: nil},
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		DeletedAt:   deletedAt,
	}

	return f
}
