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
	s.Handle("/create-todo", mc.AuthMiddleware(tc.CreateTodo))
	s.Handle("/read-todo", mc.AuthMiddleware(tc.ReadTodo))
	s.Handle("/update-todo", mc.AuthMiddleware(tc.UpdateTodo))
	s.Handle("/delete-todo", mc.AuthMiddleware(tc.DeleteTodo))
	return s
}
