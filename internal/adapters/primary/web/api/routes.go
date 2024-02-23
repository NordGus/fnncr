package api

import (
	"financo/internal/adapters/primary/web/api/app/accounts"
	"financo/internal/adapters/primary/web/api/app/applets"
	"financo/internal/adapters/primary/web/api/app/authentication"
	"financo/internal/adapters/primary/web/api/app/savingsgoals"
	"github.com/labstack/echo/v4/middleware"
)

func (a *App) setRoutes() {
	auth := authentication.New(a.AuthAPI)
	apptls := applets.New()
	accnts := accounts.New()
	svngsgls := savingsgoals.New()

	a.echo.Use(middleware.Logger())

	a.echo.StaticFS("/dist", a.assetsFS)

	a.echo.GET(authentication.LoginRoute, auth.LoginHandlerFunc)
	a.echo.POST(authentication.SignInRoute, auth.SignInHandlerFunc)
	a.echo.GET(authentication.SignOutRoute, auth.SignOutHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET("/", apptls.RootAppletHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(applets.DashboardAppletRoute, apptls.DashboardAppletHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(applets.BookAppletRoute, apptls.BookAppletHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(applets.BudgetAppletRoute, apptls.BudgetAppletHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(applets.IntelligenceAppletRoute, apptls.IntelligenceAppletHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET(accounts.CapitalAccountsRoute, accnts.CapitalHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.DebtAccountsRoute, accnts.DebtAccountsHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.ExternalAccountsRoute, accnts.ExternalAccountsHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.NewAccountRoute, accnts.NewHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.NewCapitalAccountRoute, accnts.NewCapitalAccountHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.NewDebtAccountRoute, accnts.NewDebtAccountHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.NewExternalAccountRoute, accnts.NewExternalAccountHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET(savingsgoals.SavingsGoalsRoute, svngsgls.ListSavingsGoalsHandlerFunc, auth.AuthorizeMiddleware)
}
