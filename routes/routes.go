package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/zoowen/simbarang/controllers"
	"github.com/zoowen/simbarang/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/token", controllers.GenerateToken)
			user.POST("/register", controllers.RegisterUser)
		}
		transaction := api.Group("/transaction").Use(middlewares.Auth())
		{
			transaction.POST("/add", controllers.AddTrxBarang)
			transaction.GET("/list", controllers.ListTransactions)
			transaction.GET("/list/:id", controllers.GetDetailTransaction)
			transaction.PUT("/list/:id", controllers.UpdateTransaction)
			// buku.DELETE("/list/:id", controllers.DeleteBuku)
		}

	}
	return router

}
