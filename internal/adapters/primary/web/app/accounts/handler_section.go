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
			{Title: "personal", Id: "personal", FetchURL: PersonalAccountsRoute},
			{Title: "debts", Id: "debts", FetchURL: DebtAccountsRoute},
			{Title: "external", Id: "external", FetchURL: ExternalAccountsRoute},
		},
	}

	return view.Applet(ald, p).Render(c.Request().Context(), c.Response())
}
