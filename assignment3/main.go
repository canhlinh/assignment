package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/canhlinh/assignment/docs" // docs is generated by Swag CLI, you have to import it.
)

func main() {
	docs.SwaggerInfo.Title = "Student API"
	docs.SwaggerInfo.Description = "Student API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r := gin.New()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/", func(ctx *gin.Context) {
		http.Redirect(ctx.Writer, ctx.Request, "/swagger/index.html", 302)
	})

	r.LoadHTMLGlob("./template/*")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})

	api := &Api{
		Store: NewStore("data.json"),
	}

	v1 := r.Group("/api/v1")
	v1.POST("/students", api.CreateStudent)
	v1.GET("/students", api.ListStudent)
	v1.GET("/students/:student_id", api.GetStudent)
	v1.PUT("/students/:student_id", api.EditStudent)

	r.Run()
}