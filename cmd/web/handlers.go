package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"gitlab.host-h.net/skunkworks/domain-api-dashboard/database"
	"gitlab.host-h.net/skunkworks/domain-api-dashboard/models"
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

	searchValue := c.FormValue("search")
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
