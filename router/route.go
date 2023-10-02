package router

import (
	"clean-architecture-todo-webapp-by-golang/controller"
	"net/http"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController, mc controller.IMiddleController) *http.ServeMux {
	s := http.NewServeMux()
	s.Handle("/", http.FileServer(http.Dir("public/")))
	s.HandleFunc("/login-api", uc.LogIn)
	s.HandleFunc("/logout-api", uc.LogOut)
	s.HandleFunc("/signup-api", uc.SignUp)
	s.Handle("/create-task", mc.AuthMiddleware(tc.CreateTask))
	s.Handle("/reload-task", mc.AuthMiddleware(tc.ReloadTask))
	s.Handle("/delete-task", mc.AuthMiddleware(tc.DeleteTask))
	return s
}
