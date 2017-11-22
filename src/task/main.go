package main

import (
   "fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
	
	
	"Routes"
	"Routes/middleware"
)
func main () {
//dir := http.Dir("./public")
router := mux.NewRouter()
static := mux.NewRouter()

	static.PathPrefix("/").Handler(http.FileServer(http.Dir("./public")))
	
	router.HandleFunc("/api/tasks", middleware.Auth(Routes.GetTasks)).Methods("GET")
	router.HandleFunc("/api/tasks", middleware.Auth(Routes.PostTask)).Methods("POST")
	router.HandleFunc("/api/tasks/{name}", middleware.Auth(Routes.GetOneTask)).Methods("GET")
	router.HandleFunc("/api/tasks/{name}", middleware.Auth(Routes.UpdateTask)).Methods("PUT")
	router.HandleFunc("/api/tasks/{name}", middleware.Auth(Routes.DeleteTask)).Methods("DELETE")
	router.HandleFunc("/api/token", Routes.GetToken).Methods("POST")
	router.HandleFunc("/api/user", Routes.CreateUser).Methods("POST")
	router.HandleFunc("/api/user", middleware.Auth(Routes.GetUser)).Methods("GET")
	router.HandleFunc("/api/user", middleware.Auth(Routes.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/api/user", middleware.Auth(Routes.UpdateUser)).Methods("PUT")

	http.Handle("/api/", router)
	http.Handle("/", static)
   
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func homePage (w http.ResponseWriter, r *http.Request) {
	fmt.Println("home page called")
	http.FileServer(http.Dir("./public"))
}
