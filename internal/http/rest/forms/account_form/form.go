package account_form

import (
	"financo/internal/http/rest/models/shared"
	"time"

	model "financo/internal/http/rest/models/account_model"
	"github.com/google/uuid"
)

type Form struct {
	ParentID ParentID

	Type      Type
	Currency  Currency
	Limit     Limit
	IsArchive IsArchive

	Name        Name
	Description Description
	Color       Color
	Icon        Icon

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
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
