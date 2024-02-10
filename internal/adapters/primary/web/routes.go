package web

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/accounts"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/applets"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/authentication"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/savingsgoals"
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

	a.echo.GET(accounts.NormalAccountsRoute, accnts.NormalAccountsHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.LoanAccountsRoute, accnts.LoanAccountsHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.CreditAccountsRoute, accnts.CreditAccountsHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.SavingsAccountsRoute, accnts.SavingsAccountsHandlerFunc, auth.AuthorizeMiddleware)
	a.echo.GET(accounts.ExternalAccountsRoute, accnts.ExternalAccountsHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET(savingsgoals.SavingsGoalsRoute, svngsgls.ListSavingsGoalsHandlerFunc, auth.AuthorizeMiddleware)
}
