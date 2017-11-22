package Routes

import (
	"strconv"
	"encoding/json"
	"fmt"
	"net/http"


	"db/models"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
	var email string = r.Header.Get("email")
	var user *models.User = &models.User{Email: email}
	err :=  user.GetOne()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Server error")
		return
	}
	user.Password = ""
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Body == nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Please supply email, name and password")
		return
	}
	 err := json.NewDecoder(r.Body).Decode(&user)
	 if err != nil {
      w.WriteHeader(500)
	  fmt.Fprintln(w, "Server error")
	  return
	}

	if user.Email == "" && user.Password == "" && user.Name == "" {
	    w.WriteHeader(400)
		fmt.Fprintln(w, "Please supply all required arguments")	
		return
	}
	user.Create()
    user.Password = ""
	json.NewEncoder(w).Encode(user)
}

func DeleteUser (w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("email")
	var user *models.User = &models.User{Email: email}
	err := user.Delete()
	if err != nil {
	   w.WriteHeader(500)
	  fmt.Fprintln(w, "Server error")
	  return
	}
     w.WriteHeader(204)
	  fmt.Fprintln(w, "Done")
	  return 
}

func UpdateUser (w http.ResponseWriter, r *http.Request) {
	value := r.Header.Get("id")
	id, _ := strconv.Atoi(value)
	var user *models.User = &models.User{Id: id}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
	  w.WriteHeader(500)
	  fmt.Fprintln(w, "Server error")
	  return
	}
	err = user.Update()
	if err != nil {
	  w.WriteHeader(500)
	  fmt.Fprintln(w, err)
	  return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}