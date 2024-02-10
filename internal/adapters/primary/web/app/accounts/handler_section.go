package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/savingsgoals"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/shared"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/layouts"
	"github.com/labstack/echo/v4"
)

func (h *handler) AppletHandlerFunc(c echo.Context) error {
	ald := c.Get(shared.ALDContextKey).(layouts.ApplicationLayoutData)
	p := view.Page{
		Title: "dashboard",
		Sections: []view.Section{
			{
				Title:    "normal accounts",
				Id:       "normal",
				FetchURL: NormalAccountsRoute,
			},
			{
				Title:    "credit lines",
				Id:       "credit",
				FetchURL: CreditAccountsRoute,
			},
			{
				Title:    "loans",
				Id:       "loans",
				FetchURL: LoanAccountsRoute,
			},
			{
				Title:    "savings accounts",
				Id:       "savings",
				FetchURL: SavingsAccountsRoute,
			},
			{
				Title:    "saving goals",
				Id:       "saving_goals",
				FetchURL: savingsgoals.SavingsGoalsRoute,
			},
			{
				Title:    "external accounts",
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
