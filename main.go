package main

import (
	"belajar-go/config"
	"belajar-go/controller"
	"belajar-go/models"
	"belajar-go/services"
	"belajar-go/storage"
	"log"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("IP: %s  METHOD: %s  PATH: %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {

	db := config.ConnectDB()
	db.AutoMigrate(&models.User{})
	repoUser := storage.NewPostgresRepo(db)
	serviceUser := services.NewUserService(repoUser, db)
	controllerUser := controller.NewUserController(serviceUser)

	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get database connection")
	}
	defer sqlDB.Close()
	http.HandleFunc("GET /users/{id}", controllerUser.FindByIdHandler)
	http.HandleFunc("POST /users", controllerUser.CreateUserHandler)
	http.HandleFunc("POST /users/{id}/topup", controllerUser.TopUpHandler)
	http.HandleFunc("POST /login", controllerUser.LoginHandler)

	http.ListenAndServe(":8080", Logger(http.DefaultServeMux))

}
