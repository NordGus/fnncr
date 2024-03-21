package account_form

import (
	"time"

	"financo/internal/forms/shared/form_value"
	"financo/internal/forms/shared/nullable_value"
)

type (
	Account interface {
		GetID() string
		GetParentID() string

		GetKind() string
		GetCurrency() string
		GetLimit() int64
		WasArchived() bool

		GetName() string
		GetDescription() string
		GetColor() string
		GetIcon() string

		GetInitialAmount() int64
		GetOpenedAt() time.Time
	}

	Form struct {
		ID       nullable_value.Value[string] `json:"id"`
		ParentID nullable_value.Value[string] `json:"parentID"`

		Kind      form_value.Value[string] `json:"kind"`
		Currency  form_value.Value[string] `json:"currency"`
		Limit     form_value.Value[int64]  `json:"limit"`
		IsArchive form_value.Value[bool]   `json:"isArchive"`

		Name        form_value.Value[string] `json:"name"`
		Description form_value.Value[string] `json:"description"`
		Color       form_value.Value[string] `json:"color"`
		Icon        form_value.Value[string] `json:"icon"`

		InitialAmount form_value.Value[int64]     `json:"initialAmount"`
		OpenedAt      form_value.Value[time.Time] `json:"openedAt"`

		Children []Form `json:"children"`

		IsValid bool `json:"isValid"`
	}
)

func NewEntry(raw Form, validator Validator) Form {
	f := Form{
		Kind:          form_value.New(raw.Kind.Value, validator.KindValidators()...),
		Currency:      form_value.New(raw.Currency.Value, validator.CurrencyValidators()...),
		Limit:         form_value.New(raw.Limit.Value, validator.LimitValidators()...),
		IsArchive:     form_value.New(raw.IsArchive.Value, validator.IsArchiveValidators()...),
		Name:          form_value.New(raw.Name.Value, validator.NameValidators()...),
		Description:   form_value.New(raw.Description.Value, validator.DescriptionValidators()...),
		Color:         form_value.New(raw.Description.Value, validator.ColorValidators()...),
		Icon:          form_value.New(raw.Icon.Value, validator.IconValidators()...),
		InitialAmount: form_value.New(raw.InitialAmount.Value, validator.InitialAmountValidators()...),
		OpenedAt:      form_value.New(raw.OpenedAt.Value, validator.OpenedAtValidators()...),
	}

	f.Validate()

	return f
}

func NewChildEntry(raw Form, validator Validator) Form {
	f := Form{
		ParentID:      nullable_value.New(raw.ParentID.Value, validator.ParentIDValidators()...),
		Kind:          form_value.New(raw.Kind.Value, validator.KindValidators()...),
		Currency:      form_value.New(raw.Currency.Value, validator.CurrencyValidators()...),
		Limit:         form_value.New(raw.Limit.Value, validator.LimitValidators()...),
		IsArchive:     form_value.New(raw.IsArchive.Value, validator.IsArchiveValidators()...),
		Name:          form_value.New(raw.Name.Value, validator.NameValidators()...),
		Description:   form_value.New(raw.Description.Value, validator.DescriptionValidators()...),
		Color:         form_value.New(raw.Description.Value, validator.ColorValidators()...),
		Icon:          form_value.New(raw.Icon.Value, validator.IconValidators()...),
		InitialAmount: form_value.New(raw.InitialAmount.Value, validator.InitialAmountValidators()...),
		OpenedAt:      form_value.New(raw.OpenedAt.Value, validator.OpenedAtValidators()...),
	}

	f.Validate()

	return f
}

func New(account Account, children ...Form) Form {
	f := Form{
		ID:            nullable_value.New(account.GetID()),
		Kind:          form_value.New(account.GetKind()),
		Currency:      form_value.New(account.GetCurrency()),
		Limit:         form_value.New(account.GetLimit()),
		IsArchive:     form_value.New(account.WasArchived()),
		Name:          form_value.New(account.GetName()),
		Description:   form_value.New(account.GetDescription()),
		Color:         form_value.New(account.GetColor()),
		Icon:          form_value.New(account.GetIcon()),
		InitialAmount: form_value.New(account.GetInitialAmount()),
		OpenedAt:      form_value.New(account.GetOpenedAt()),
		Children:      children,
		IsValid:       true,
	}

	return f
}

func NewChild(account Account) Form {
	f := Form{
		ID:            nullable_value.New(account.GetID()),
		ParentID:      nullable_value.New(account.GetParentID()),
		Kind:          form_value.New(account.GetKind()),
		Currency:      form_value.New(account.GetCurrency()),
		Limit:         form_value.New(account.GetLimit()),
		IsArchive:     form_value.New(account.WasArchived()),
		Name:          form_value.New(account.GetName()),
		Description:   form_value.New(account.GetDescription()),
		Color:         form_value.New(account.GetColor()),
		Icon:          form_value.New(account.GetIcon()),
		InitialAmount: form_value.New(account.GetInitialAmount()),
		OpenedAt:      form_value.New(account.GetOpenedAt()),
		IsValid:       true,
	}

	return f
}

func (f *Form) Validate() {
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
}
