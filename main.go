package main

import (
	"belajar-go/config"
	"belajar-go/controller"
	"belajar-go/models"
	"belajar-go/services"
	"belajar-go/storage"
	"net/http"
)

func main() {

	db := config.ConnectDB()
	db.AutoMigrate(&models.User{})
	repoUser := storage.NewPostgresRepo(db)
	serviceUser := services.NewUserService(repoUser)
	controllerUser := controller.NewUserController(serviceUser)

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	defer sqlDB.Close()
	http.HandleFunc("GET /users/{id}", controllerUser.FindById)
	http.HandleFunc("POST /users", controllerUser.CreateUser)

	http.ListenAndServe(":8080", nil)

}
