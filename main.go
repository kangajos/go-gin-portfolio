package main

import (
	"os"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kangajos/go-gin-portfolio.git/routes"
)

func main() {
	godotenv.Load(".env")

	app := gin.Default()
	app.Static("/assets/", "./assets")
	app.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "templates",
		Extension:    ".html",
		Master:       "layouts/web",
		Partials:     []string{"partials/header", "partials/navbar", "partials/footer"},
		DisableCache: true,
	})
	routes.Portfolio(app)
	app.Run(":" + os.Getenv("APP_PORT"))
}
