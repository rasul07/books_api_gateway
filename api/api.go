package api

import (
	"github.com/rasul07/books_api_gateway/config"
	"github.com/rasul07/books_api_gateway/pkg/logger"
	"github.com/rasul07/books_api_gateway/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"


	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// _ "github.com/rasul07/books_api_gateway/api/docs"
	v1 "github.com/rasul07/books_api_gateway/api/handlers/v1"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	cfg := cors.DefaultConfig()

	cfg.AllowHeaders = append(cfg.AllowHeaders, "*")
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true

	router.Use(cors.New(cfg))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:     opt.Log,
		Cfg:     opt.Cfg,
		Services: opt.Services,
	})

	apiV1 := router.Group("/v1")

	{
		// Book
		apiV1.POST("/book", handlerV1.CreateBook)
		apiV1.GET("/book", handlerV1.GetBooks)
		apiV1.GET("/book/:book_id", handlerV1.GetBookById)
		apiV1.PUT("/book/:book_id", handlerV1.UpdateBook)
		apiV1.DELETE("/book/:book_id", handlerV1.Delete)

		// Category 
		apiV1.POST("/category", handlerV1.Create)
		apiV1.GET("/category", handlerV1.GetCategories)
		apiV1.GET("/category/:category_id", handlerV1.GetCategoryById)
		apiV1.PUT("/category", handlerV1.UpdateCategory)
		apiV1.DELETE("/category/:category_id", handlerV1.DeleteCategory)

	}

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
