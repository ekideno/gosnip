package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"time"

	"github.com/ekideno/gosnip/internal/database"
	"github.com/ekideno/gosnip/internal/handler"
	"github.com/ekideno/gosnip/internal/repository"
	"github.com/ekideno/gosnip/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://gosnip"}, // Разрешённый origin
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlService)

	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.POST("/random", urlHandler.CreateRandom)
	r.GET("/:slug", urlHandler.GetBySlug)

	err = r.Run(":9000")
	if err != nil {
		log.Fatal(err)
	}
}
