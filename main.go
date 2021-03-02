package main

import (
	"lyra/handler"
	"lyra/request"
	"lyra/subscribe"
	"lyra/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/lyra-music?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	userRepository := user.NewRepository(db)
	subscribeRepository := subscribe.NewRepository(db)
	requestRepository := request.NewRepository(db)

	userSerice := user.NewService(userRepository)
	subscribeService := subscribe.NewService(subscribeRepository)
	requestService := request.NewService(requestRepository)

	userHandler := handler.UserHandler(userSerice)
	subscribeHandler := handler.SubscribeHandler(subscribeService)
	requestHandler := handler.RequestHandler(requestService)

	router := gin.Default()
	router.POST("/register-admin", userHandler.Register)
	router.POST("/login-admin", userHandler.Login)
	router.POST("/user-token", userHandler.TokenUser)

	//subscribe
	router.POST("/subscriber", subscribeHandler.Create)
	router.GET("/subscriber", subscribeHandler.FindAll)
	router.DELETE("/subscriber/:id", subscribeHandler.Delete)

	//request music
	router.POST("/request-music", requestHandler.Create)
	router.GET("/request-music", requestHandler.Find)
	router.DELETE("/request-music/:id", requestHandler.Delete)

	router.Run(":2020")
}
