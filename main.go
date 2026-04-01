package main

import (
	"belajar-go/config"
	"belajar-go/controller"
	"belajar-go/models"
	"belajar-go/services"
	"belajar-go/storage"
	"net/http"
	"belajar-go/logger"
	"belajar-go/middleware"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	logger.InitLogger()
	defer logger.Log.Sync()


	db := config.ConnectDB()
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	defer sqlDB.Close()


	db.AutoMigrate(&models.User{})
	repoUser := storage.NewPostgresRepo(db)
	serviceUser := services.NewIUserService(repoUser, db)
	controllerUser := controller.NewUserController(serviceUser)

	mux := http.NewServeMux()
	handler := middleware.RecoveryMiddleware(middleware.LoggingMiddleware(mux))

	mux.Handle("GET /users", middleware.TokenMiddleware(http.HandlerFunc(controllerUser.FindByIdHandler)))
	mux.HandleFunc("POST /users", controllerUser.CreateUserHandler)
	mux.Handle("POST /users/topup", middleware.TokenMiddleware(http.HandlerFunc(controllerUser.TopUpHandler)))
	mux.HandleFunc("POST /login", controllerUser.LoginHandler)

	http.ListenAndServe(":8080", handler)
}
