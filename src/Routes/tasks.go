package Routes

import (
	"strconv"
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	

	"db/models"
)

func GetTasks (w http.ResponseWriter, r *http.Request) {
	 value := r.Header.Get("id")
	 id, _ := strconv.Atoi(value)
   var tasks *models.Tasks = &models.Tasks{User_id: id}
   
   taskArray, err := tasks.GetAll()
   if err != nil {
	   w.WriteHeader(500)
		fmt.Fprintln(w, "Server error")
		return
   }

   json.NewEncoder(w).Encode(taskArray)
}

func PostTask (w http.ResponseWriter, r *http.Request) {
	value := r.Header.Get("id")
	id, _ := strconv.Atoi(value)
	if r.Body == nil {
	  w.WriteHeader(400)
		fmt.Fprintln(w, "Please supply a task")
		return
	}
	var task *models.Tasks = &models.Tasks{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
	  w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}

	task.User_id = id
	fmt.Println(task)
	 
	err = task.Create()
   if err != nil {
	   w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}
	 json.NewEncoder(w).Encode(task)
}

func GetOneTask (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	value := r.Header.Get("id")
	id, _ := strconv.Atoi(value)

	var task *models.Tasks = &models.Tasks{}
	task.Title = name
	task.User_id = id
	err := task.GetOne()
	if err != nil {
	   w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}
    json.NewEncoder(w).Encode(task)
}

func UpdateTask (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	value := r.Header.Get("id")
	id, _ := strconv.Atoi(value)
	
  if r.Body == nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Supply done parameter")
		return
	}
	var task *models.Tasks = &models.Tasks{}
	task.Title = name
	task.User_id = id
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
	  w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}
	err = task.Update()
	if err != nil {
	  w.WriteHeader(500)
		fmt.Fprintln(w, err)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func DeleteTask (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	title := params["name"]
	value := r.Header.Get("id")
	id, _ := strconv.Atoi(value)

	var task *models.Tasks = &models.Tasks{}
	task.Title = title
	task.User_id = id
	err := task.Delete()
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Server error")
		return
	}
	fmt.Fprintln(w, "Done")
}