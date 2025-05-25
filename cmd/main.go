package main

import (
	"github.com/ekideno/gosnip/internal/database"
	"github.com/ekideno/gosnip/internal/handler"
	"github.com/ekideno/gosnip/internal/repository"
	"github.com/ekideno/gosnip/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlService)

	api := r.Group("/api")
	api.POST("/random", urlHandler.CreateRandom)
	api.GET("/:slug", urlHandler.GetBySlug)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
