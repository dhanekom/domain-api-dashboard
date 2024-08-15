package main

import (
	"context"
	"strings"

	"github.com/labstack/echo/v4"
	"gitlab.host-h.net/skunkworks/domain-api-dashboard/internal/database"
	"gitlab.host-h.net/skunkworks/domain-api-dashboard/internal/models"
)

type HandlerRepo struct {
	db database.DBRepo
}

func newHandlerRepo(db database.DBRepo) *HandlerRepo {
	return &HandlerRepo{
		db: db,
	}
}

func (r *HandlerRepo) DashboardHandler(c echo.Context) error {
	return c.Render(200, "dashboard.html", nil)
}

func (r *HandlerRepo) DashboardSearchResultsHandler(c echo.Context) error {
	data := models.NewDashboardData()

	searchValue := strings.Trim(c.FormValue("search"), " ")

	if len(searchValue) < 3 {
		return c.Render(422, "notification", models.Notification{Message: "A valid search value is required"})
	}

	searchByCode := c.FormValue("search-by")
	searchByType := models.DomainSearchByTypeFromCode(searchByCode)

	var err error
	data.Domains, err = r.db.GetDomains(context.Background(), searchByType, searchValue)
	if err != nil {
		return err
	}

	// sort.Slice(data.Domains, func(i, j int) bool {
	// 	return data.Domains[i].DomainName < data.Domains[j].DomainName
	// })

	return c.Render(200, "dashboard-search-result", data)
}
