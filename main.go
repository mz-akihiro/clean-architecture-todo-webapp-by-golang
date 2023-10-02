package main

import (
	"clean-architecture-todo-webapp-by-golang/controller"
	"clean-architecture-todo-webapp-by-golang/db"
	"clean-architecture-todo-webapp-by-golang/repository"
	"clean-architecture-todo-webapp-by-golang/router"
	"clean-architecture-todo-webapp-by-golang/usecase"
	"log"
	"net/http"
)

func main() {
	db := db.Newdb()
	defer db.Close()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	middleController := controller.NewMiddleController()
	server := router.NewRouter(userController, taskController, middleController)
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
