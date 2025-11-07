package main

import (
	"go-gin-fw/database"
	"go-gin-fw/routes"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	tmpl := template.Must(template.ParseGlob("views/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("views/user/*.html"))

	r.SetHTMLTemplate(tmpl)

	routes.SetupRoutes(r)

	r.Run(":8080")
}
