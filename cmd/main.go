package main

import (
	"github.com/gin-gonic/gin"
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	// Camada useCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProductList)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("product/:productId", ProductController.GetProductById)

	server.Run(":8000")
}
