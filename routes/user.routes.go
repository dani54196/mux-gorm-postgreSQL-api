package routes

import "net/http"

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UserS"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PostUser"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
