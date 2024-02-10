package applets

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/accounts"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/savingsgoals"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) DashboardAppletHandlerFunc(c echo.Context) error {
	sections := []view.DashboardSection{
		{
			Title:    "normal accounts",
			Id:       "normal",
			FetchURL: accounts.NormalAccountsRoute,
		},
		{
			Title:    "credit lines",
			Id:       "credit",
			FetchURL: accounts.CreditAccountsRoute,
		},
		{
			Title:    "loans",
			Id:       "loans",
			FetchURL: accounts.LoanAccountsRoute,
		},
		{
			Title:    "savings accounts",
			Id:       "savings",
			FetchURL: accounts.SavingsAccountsRoute,
		},
		{
			Title:    "saving goals",
			Id:       "saving_goals",
			FetchURL: savingsgoals.SavingsGoalsRoute,
		},
		{
			Title:    "external accounts",
			Id:       "external",
			FetchURL: accounts.ExternalAccountsRoute,
		},
	}

	return view.Dashboard(
		layoutData(
			getUser(c),
			"fnncr | dashboard",
			dashboard,
		),
		sections,
	).Render(c.Request().Context(), c.Response())
}
