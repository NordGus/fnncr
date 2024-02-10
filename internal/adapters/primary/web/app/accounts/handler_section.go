package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/shared"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/layouts"
	"github.com/labstack/echo/v4"
)

func (h *handler) AppletHandlerFunc(c echo.Context) error {
	ald := c.Get(shared.ALDContextKey).(layouts.ApplicationLayoutData)
	p := view.Page{
		Title: "accounts",
		Sections: []view.Section{
			{
				Title:    "personal",
				Id:       "personal",
				FetchURL: NormalAccountsRoute,
			},
			{
				Title:    "loans",
				Id:       "loans",
				FetchURL: LoanAccountsRoute,
			},
			{
				Title:    "credit",
				Id:       "credit",
				FetchURL: CreditAccountsRoute,
			},
			{
				Title:    "savings",
				Id:       "savings",
				FetchURL: SavingsAccountsRoute,
			},
			{
				Title:    "saving goals",
				Id:       "saving_goals",
				FetchURL: SavingsGoalsAccountsRoute,
			},
			{
				Title:    "external",
				Id:       "external",
				FetchURL: ExternalAccountsRoute,
			},
		},
	}

	ald.Title = "fnncr | accounts"

	for i := 0; i < len(ald.NavItems); i++ {
		if ald.NavItems[i].Name == "accounts" {
			ald.NavItems[i].IsActive = true

			break
		}
	}

	return view.Applet(ald, p).Render(c.Request().Context(), c.Response())
}
