package main

import (
	"log"
	"shopApi/db"
	"shopApi/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключаемся к базе данных
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	router := gin.Default()

	// Роуты для продуктов
	router.GET("/products", handlers.GetProducts(db))
	router.GET("/products/:id", handlers.GetProduct(db))
	router.POST("/products", handlers.CreateProduct(db))
	router.PUT("/products/:id", handlers.UpdateProduct(db))
	router.DELETE("/products/:id", handlers.DeleteProduct(db))

	// Роуты для корзины
	router.GET("/carts/:id", handlers.GetCart(db))
	router.POST("/carts/:userId", handlers.AddToCart(db))
	router.DELETE("/carts/:userId/:productId", handlers.RemoveFromCart(db))

	// Роуты для избранного
	router.GET("/favorites/:id", handlers.GetFavorites(db))
	router.POST("/favorites/:userId", handlers.AddToFavorites(db))
	router.DELETE("/favorites/:userId/:productId", handlers.RemoveFromFavorites(db))

	// Роуты для заказов
	router.GET("/orders/:id", handlers.GetOrders(db))
	router.POST("/orders/:id", handlers.CreateOrder(db))

	// Запуск сервера
	router.Run(":8080")
}
