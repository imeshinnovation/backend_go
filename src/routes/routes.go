package routes

import (
	"netix/src/controllers"
	"netix/src/middleware"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(db *mongo.Database) *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.AuthMiddleware)

	userController := controllers.NewUserController(db)

	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/login", userController.LoginUser).Methods("POST")

	// Rutas protegidas
	router.HandleFunc("/users/{id}", userController.GetUser).Methods("GET")
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	return router
}
