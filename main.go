package main

import (
	"belajar-go/config"
	"belajar-go/controller/handlerTransactions"
	"belajar-go/controller/handlerUser"
	"belajar-go/logger"
	"belajar-go/middleware"
	"belajar-go/models"
	"belajar-go/services/email"
	"belajar-go/services/transactions"
	"belajar-go/services/user"
	"belajar-go/storage"
	"belajar-go/utils"
	"github.com/joho/godotenv"
	"net/http"
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
	utils.StartWorker()

	db.AutoMigrate(&models.User{}, &models.Transaction{})

	repoUser := storage.NewPostgresRepo(db)
	repoTransaction := storage.NewTransactionRepo(db)
	EmailService := email.NewEmailService()

	serviceUser := user.NewIUserService(repoUser, db, EmailService)
	controllerUser := handlerUser.NewUserController(serviceUser)

	ServiceTransaction := transactions.NewITransactionService(repoTransaction, repoUser, db, EmailService)
	controllerTransaction := handlertransactions.NewControllerTransaction(ServiceTransaction)

	mux := http.NewServeMux()
	handler := middleware.RecoveryMiddleware(middleware.LoggingMiddleware(mux))

	mux.Handle("GET /users", middleware.TokenMiddleware(http.HandlerFunc(controllerUser.FindByIdHandler)))
	mux.HandleFunc("POST /users", controllerUser.CreateUserHandler)
	mux.Handle("POST /users/topup", middleware.TokenMiddleware(http.HandlerFunc(controllerTransaction.TopUpHandler)))
	mux.HandleFunc("POST /login", controllerUser.LoginHandler)
	mux.Handle("GET /transactions", middleware.TokenMiddleware(http.HandlerFunc(controllerTransaction.GetHistoryTransactionHandler)))
	mux.HandleFunc("POST /notifications-midtrans", controllerTransaction.HandleMidtransNotification)

	http.ListenAndServe(":8080", handler)
}
