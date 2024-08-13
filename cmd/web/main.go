package main

import (
	"html/template"
	"io"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.host-h.net/skunkworks/domain-api-dashboard/models"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type DashboardData struct {
	Domains []models.Domain
}

func newDashboardData() *DashboardData {
	return &DashboardData{
		Domains: []models.Domain{},
	}
}

func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	e.GET("/", dashboardHandler)
	e.POST("/search", dashboardSearchResultsHandler)

	e.Logger.Fatal(e.Start("127.0.0.1:42069"))
}

func dashboardHandler(c echo.Context) error {
	return c.Render(200, "dashboard", nil)
}

func dashboardSearchResultsHandler(c echo.Context) error {
	data := newDashboardData()

	searchValue := c.FormValue("search")
	// if strings.Trim(searchValue, " ") == "" {
	// 	return c.Render(200, "dashboard-search-result", data)
	// }

	domains := []models.Domain{
		{DomainName: "virtua.co.za", AccountNo: "A0100591103", ClientNo: "C0100405403", BillForReg: false, Status: "Ready"},
		{DomainName: "moneyforyou.co.za", AccountNo: "C0301767006", ClientNo: "A0100591303", BillForReg: false, Status: "Ready"},
		{DomainName: "computerknights.co.za", AccountNo: "A0100592203", ClientNo: "C0100406203", BillForReg: false, Status: "Ready"},
		{DomainName: "lexi.co.za", AccountNo: "A0100592203", ClientNo: "C0100406203", BillForReg: false, Status: "Deleted"},
	}

	data.Domains = filter[models.Domain](domains, func(d models.Domain) bool {
		return strings.HasPrefix(d.DomainName, searchValue)
	})

	return c.Render(200, "dashboard-search-result", data)
}
