package main

import (
	"context"
	"html/template"
	"io"
	"log"
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.host-h.net/skunkworks/domain-api-dashboard/database"
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

type HandlerRepo struct {
	db database.DBRepo
}

func newHandlerRepo(db database.DBRepo) *HandlerRepo {
	return &HandlerRepo{
		db: db,
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONN_STRING")
	db, err := database.ConnectToDB(dsn)
	if err != nil {
		log.Fatalf("unable to connect to db: %s", err)
	}
	defer db.Close()

	dbRepo := database.NewMysqlDBRepo(db)

	handlerRepo := newHandlerRepo(dbRepo)

	e := echo.New()

	e.Static("/public", "public")
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	e.GET("/", handlerRepo.DashboardHandler)
	e.POST("/search", handlerRepo.DashboardSearchResultsHandler)

	e.Logger.Fatal(e.Start("127.0.0.1:42069"))
}

func (r *HandlerRepo) DashboardHandler(c echo.Context) error {
	return c.Render(200, "dashboard", nil)
}

func (r *HandlerRepo) DashboardSearchResultsHandler(c echo.Context) error {
	data := newDashboardData()

	searchValue := c.FormValue("search")
	// if strings.Trim(searchValue, " ") == "" {
	// 	return c.Render(200, "dashboard-search-result", data)
	// }

	// domains := []models.Domain{
	// 	{DomainName: "virtua.co.za", AccountNo: "A0100591103", ClientNo: "C0100405403", BillForReg: false, Status: "Ready"},
	// 	{DomainName: "moneyforyou.co.za", AccountNo: "C0301767006", ClientNo: "A0100591303", BillForReg: false, Status: "Ready"},
	// 	{DomainName: "computerknights.co.za", AccountNo: "A0100592203", ClientNo: "C0100406203", BillForReg: false, Status: "Ready"},
	// 	{DomainName: "lexi.co.za", AccountNo: "A0100592203", ClientNo: "C0100406203", BillForReg: false, Status: "Deleted"},
	// }

	var err error
	data.Domains, err = r.db.GetDomains(context.Background(), models.DomainSearchByDomainName, searchValue)
	if err != nil {
		return err
	}

	sort.Slice(data.Domains, func(i, j int) bool {
		return data.Domains[i].DomainName < data.Domains[j].DomainName
	})

	// data.Domains = filter[models.Domain](domains, func(d models.Domain) bool {
	// 	return strings.HasPrefix(d.DomainName, searchValue)
	// })

	return c.Render(200, "dashboard-search-result", data)
}
