package account_form

import (
	"time"

	model "financo/internal/http/rest/models/account_model"
	"financo/internal/http/rest/models/shared"
	"github.com/google/uuid"
)

type Form struct {
	ParentID ParentID `json:"parentID"`

	Type      Type      `json:"type"`
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
	t model.Type,
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
) Form {
	f := Form{
		Type:        Type{Value: t, Errors: nil},
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
	t model.Type,
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
) Form {
	f := Form{
		ParentID:    ParentID{Value: parentID, Errors: nil},
		Type:        Type{Value: t, Errors: nil},
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
