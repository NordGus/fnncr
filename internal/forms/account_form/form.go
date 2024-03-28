package account_form

import (
	"time"

	"financo/internal/forms/shared/form_value"
	"financo/internal/forms/shared/nullable_value"
)

type (
	Account interface {
		ID() string
		ParentID() string

		Kind() string
		Currency() string
		Limit() int64
		Archived() bool

		GetName() string
		Description() string
		Color() string
		Icon() string

		InitialAmount() int64
		OpenedAt() time.Time
	}

	Form struct {
		ID       form_value.Value[string]     `json:"id"`
		ParentID nullable_value.Value[string] `json:"parentID"`

		Kind      form_value.Value[string] `json:"kind"`
		Currency  form_value.Value[string] `json:"currency"`
		Limit     form_value.Value[int64]  `json:"limit"`
		IsArchive form_value.Value[bool]   `json:"isArchive"`

		Name        form_value.Value[string] `json:"name"`
		Description form_value.Value[string] `json:"description"`
		Color       form_value.Value[string] `json:"color"`
		Icon        form_value.Value[string] `json:"icon"`

		InitialAmount nullable_value.Value[int64]     `json:"initialAmount"`
		OpenedAt      nullable_value.Value[time.Time] `json:"openedAt"`

		Children []Form `json:"children"`

		IsValid bool `json:"isValid"`

		initialized bool
	}
)

func NewEntry(raw Form) Form {
	f := Form{
		ID:            form_value.New(raw.ID.Value),
		ParentID:      nullable_value.New("", false),
		Kind:          form_value.New(raw.Kind.Value),
		Currency:      form_value.New(raw.Currency.Value),
		Limit:         form_value.New(raw.Limit.Value),
		IsArchive:     form_value.New(raw.IsArchive.Value),
		Name:          form_value.New(raw.Name.Value),
		Description:   form_value.New(raw.Description.Value),
		Color:         form_value.New(raw.Description.Value),
		Icon:          form_value.New(raw.Icon.Value),
		InitialAmount: nullable_value.New(raw.InitialAmount.Value, true),
		OpenedAt:      nullable_value.New(raw.OpenedAt.Value, raw.OpenedAt.Value.IsZero()),
		initialized:   true,
	}

	f.Valid()

	return f
}

func NewChildEntry(raw Form) Form {
	f := Form{
		ID:            form_value.New(raw.ID.Value),
		ParentID:      nullable_value.New(raw.ParentID.Value, true),
		Kind:          form_value.New(raw.Kind.Value),
		Currency:      form_value.New(raw.Currency.Value),
		Limit:         form_value.New(raw.Limit.Value),
		IsArchive:     form_value.New(raw.IsArchive.Value),
		Name:          form_value.New(raw.Name.Value),
		Description:   form_value.New(raw.Description.Value),
		Color:         form_value.New(raw.Description.Value),
		Icon:          form_value.New(raw.Icon.Value),
		InitialAmount: nullable_value.New(raw.InitialAmount.Value, true),
		OpenedAt:      nullable_value.New(raw.OpenedAt.Value, raw.OpenedAt.Value.IsZero()),
		initialized:   true,
	}

	f.Valid()

	return f
}

func New(account Account, children ...Form) Form {
	f := Form{
		ID:            form_value.New(account.ID()),
		ParentID:      nullable_value.New("", false),
		Kind:          form_value.New(account.Kind()),
		Currency:      form_value.New(account.Currency()),
		Limit:         form_value.New(account.Limit()),
		IsArchive:     form_value.New(account.Archived()),
		Name:          form_value.New(account.GetName()),
		Description:   form_value.New(account.Description()),
		Color:         form_value.New(account.Color()),
		Icon:          form_value.New(account.Icon()),
		InitialAmount: nullable_value.New(account.InitialAmount(), true),
		OpenedAt:      nullable_value.New(account.OpenedAt(), account.OpenedAt().IsZero()),
		Children:      children,
		initialized:   true,
	}

	f.Valid()

	return f
}

func NewChild(account Account) Form {
	f := Form{
		ID:            form_value.New(account.ID()),
		ParentID:      nullable_value.New(account.ParentID(), true),
		Kind:          form_value.New(account.Kind()),
		Currency:      form_value.New(account.Currency()),
		Limit:         form_value.New(account.Limit()),
		IsArchive:     form_value.New(account.Archived()),
		Name:          form_value.New(account.GetName()),
		Description:   form_value.New(account.Description()),
		Color:         form_value.New(account.Color()),
		Icon:          form_value.New(account.Icon()),
		InitialAmount: nullable_value.New(account.InitialAmount(), true),
		OpenedAt:      nullable_value.New(account.OpenedAt(), account.OpenedAt().IsZero()),
		initialized:   true,
	}

	f.Valid()

	return f
}

func (f *Form) Initialized() bool {
	return f.initialized
}

func (f *Form) Valid() bool {
	f.ID.Validate()
	f.ParentID.Validate()
	f.Kind.Validate()
	f.Currency.Validate()
	f.Limit.Validate()
	f.IsArchive.Validate()
	f.Name.Validate()
	f.Description.Validate()
	f.Color.Validate()
	f.Icon.Validate()
	f.InitialAmount.Validate()
	f.OpenedAt.Validate()

	f.IsValid = f.ID.Valid() &&
		f.ParentID.Valid() &&
		f.Kind.Valid() &&
		f.Currency.Valid() &&
		f.Limit.Valid() &&
		f.IsArchive.Valid() &&
		f.Name.Valid() &&
		f.Description.Valid() &&
		f.Color.Valid() &&
		f.Icon.Valid() &&
		f.InitialAmount.Valid() &&
		f.OpenedAt.Valid()

	return f.IsValid
}
