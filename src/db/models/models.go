package models 

import (
//	"fmt"
	"log"
	"database/sql"
	"time"

	_"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)


type Tasks struct {
	Id int
	Title string
	Done bool
	User_id int
	Created_at time.Time
}
type User struct {
	Id int
	Name string
	Password string
	Email string
	Tasks []Tasks
}

var Db *sql.DB

func init() {
	var err error

	Db, err = sql.Open("postgres", "user=postgres dbname=go_app sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (task *Tasks) Create() error {
query := "insert into tasks (title, user_id) values ($1, $2) returning id, done"

stmt, err := Db.Prepare(query)

if err != nil {
    panic(err)
}

defer stmt.Close()
err = stmt.QueryRow(task.Title, task.User_id).Scan(&task.Id, &task.Done)

if err != nil {
	return err
}
return nil
}


func (task *Tasks) GetAll() ([]Tasks,  error) {
	query := "select * from tasks where user_id = $1"

	rows, err := Db.Query(query, task.User_id)
	
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	
	var tasks []Tasks

	for rows.Next() {
		var task Tasks 
		if err := rows.Scan(&task.Id, &task.Title, &task.Done, &task.User_id, &task.Created_at); err != nil {
		     return nil, err
		} 
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (task *Tasks) GetOne() error {
	query := "select * from Tasks where user_id = $1 and title = $2 "

	row, err := Db.Query(query, task.User_id, task.Title)

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer row.Close()
    for row.Next() {
		if err := row.Scan(&task.Id, &task.Title, &task.Done, &task.User_id, &task.Created_at); err != nil {
			//log.Fatal(err)
			return err
		}
	}
	return nil
}

func (task *Tasks) Update()  error{
	query := "update tasks set done = $1 where user_id = $2 and title = $3"
	
   _, err := Db.Exec(query, task.Done, task.User_id, task.Title)

   if err != nil {
	   log.Fatal(err)
	   return err
   }
    return nil
}

func (task *Tasks) Delete()  error {
	query := "delete from tasks where user_id = $1 and title = $2"

	_, err := Db.Exec(query, task.User_id, task.Title)

	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

func hashPassword(password string) (hash string, err error) {
	bytes, err:= bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes), err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (user *User) Create()  error {
	var password string = user.Password
    hash, err := hashPassword(password)
	
	
	query := "insert into users (name, email, password) values ($1, $2, $3) returning id"
	stmt, err := Db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	 err = stmt.QueryRow(user.Name, user.Email, hash).Scan(&user.Id)
	 return nil
}

func (user *User) GetOne()  error {
	query := "select id, name, password from users where email = $1"
	//fmt.Println(user)

	err := Db.QueryRow(query, user.Email).Scan(&user.Id, &user.Name, &user.Password)
	if err != nil {
		return err
	}
	id := user.Id
	
	var tasks []Tasks
	var task *Tasks = &Tasks{User_id: id}

	tasks, err = task.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	user.Tasks = tasks
    return nil
}

func (user *User) Delete()  error {
	query := "delete from users where email = $1"
	_, err := Db.Exec(query, user.Email)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (user *User) Update()  error {
 query := "update users set name = $1, email = $2 where id = $3"
 _, err := Db.Exec(query, user.Name, user.Email, user.Id)

 if err != nil {
	 //log.Fatal(err)
	 return err
 }
 return nil
}
