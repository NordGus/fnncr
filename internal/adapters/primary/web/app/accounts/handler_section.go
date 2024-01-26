package accounts

import (
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) AppletHandlerFunc(c echo.Context) error {
	p := view.Page{
		Title: "Accounts",
		Sections: []view.Section{
			{Title: "Personal", Id: "personal", FetchURL: PersonalAccountsRoute},
			{Title: "Debts", Id: "debts", FetchURL: DebtAccountsRoute},
			{Title: "External", Id: "external", FetchURL: ExternalAccountsRoute},
		},
	}

	return view.Applet(p).Render(c.Request().Context(), c.Response())
}
