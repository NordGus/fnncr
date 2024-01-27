package web

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/accounts"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/assets"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/authentication"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/budget"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/summary"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/transactions"
	views "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/application"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/components"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/layouts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *App) setRoutes() {
	auth := authentication.New(a.AuthAPI)
	accnts := accounts.New()
	trnsctns := transactions.New()
	bdgt := budget.New()
	assts := assets.New()
	smmry := summary.New()

	a.echo.Use(middleware.Logger())

	a.echo.StaticFS("/dist", a.assetsFS)

	a.echo.GET(authentication.LoginRoute, auth.LoginHandlerFunc)
	a.echo.POST(authentication.SignInRoute, auth.SignInHandlerFunc)
	a.echo.GET(authentication.SignOutRoute, auth.SignOutHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET(accounts.AppletRoute, accnts.AppletHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET(transactions.AppletRoute, trnsctns.AppletHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET(budget.AppletRoute, bdgt.AppletHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET(assets.AppletRoute, assts.AppletHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET(summary.AppletRoute, smmry.AppletHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET("/", func(c echo.Context) error {
		ald := layouts.ApplicationLayoutData{
			Title: "fnncr",
			UserOptionNave: components.NavItemWithDropdownData{
				Name: c.Get(authentication.CurrentUserCtxKey).(models.User).Username,
				Options: []components.NavItemData{
					{Name: "Sign Out", Route: authentication.SignOutRoute},
				},
			},
			NavItems: []components.NavItemData{
				{Name: "accounts", Route: accounts.AppletRoute},
				{Name: "transactions", Route: transactions.AppletRoute},
				{Name: "budget", Route: budget.AppletRoute},
				{Name: "assets", Route: assets.AppletRoute},
				{Name: "summary", Route: summary.AppletRoute},
			},
		}

		return views.Root(ald).Render(c.Request().Context(), c.Response())
	}, auth.AuthorizeMiddleware)
}
