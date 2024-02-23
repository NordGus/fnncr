package applets

import (
	auth "financo/internal/adapters/primary/web/api/app/authentication"
	"financo/internal/adapters/primary/web/api/app/models"
	"financo/internal/adapters/primary/web/api/app/views/components"
	"financo/internal/adapters/primary/web/api/app/views/layouts"
	"github.com/labstack/echo/v4"
)

const (
	// Routes

	DashboardAppletRoute    = "/dashboard"
	BookAppletRoute         = "/book"
	BudgetAppletRoute       = "/budget"
	IntelligenceAppletRoute = "/intelligence"

	// NavItemDataNames

	root         = "root"
	dashboard    = "dashboard"
	book         = "book"
	budget       = "budget"
	intelligence = "intelligence"
)

type Handler interface {
	RootAppletHandlerFunc(c echo.Context) error
	DashboardAppletHandlerFunc(c echo.Context) error
	BookAppletHandlerFunc(c echo.Context) error
	BudgetAppletHandlerFunc(c echo.Context) error
	IntelligenceAppletHandlerFunc(c echo.Context) error
}

type handler struct {
}

func New() Handler {
	return &handler{}
}

func getUser(c echo.Context) models.User {
	return c.Get(auth.CurrentUserCtxKey).(models.User)
}

func layoutData(user models.User, title string, applet string) layouts.ApplicationLayoutData {
	return layouts.ApplicationLayoutData{
		Title: title,
		UserOptionNave: components.NavItemWithDropdownData{
			Name: user.Username,
			Options: []components.NavItemData{
				{Name: "sign out", Route: auth.SignOutRoute},
			},
		},
		NavItems: []components.NavItemData{
			{Name: dashboard, Route: DashboardAppletRoute, IsActive: dashboard == applet},
			{Name: book, Route: BookAppletRoute, IsActive: book == applet},
			{Name: budget, Route: BudgetAppletRoute, IsActive: budget == applet},
			{Name: intelligence, Route: IntelligenceAppletRoute, IsActive: intelligence == applet},
		},
	}
}
