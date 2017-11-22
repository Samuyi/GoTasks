package Routes 

import (
	"fmt"
	"time"
	"encoding/json"
	"net/http"

	 "github.com/dgrijalva/jwt-go"
	 "db/models"
)

type credentials struct {
	Email string 
	Password string
	Name string
}

type Token struct {
	Token string `json:"token,omitempty"`
}
var Secret = []byte("samuyi902210")

func GetToken(w http.ResponseWriter, r *http.Request) {
	

	//token := jwt.New(jwt.SigningMethodHS256)
	
	var person credentials

	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, err)
		return
	}
	var user *models.User = &models.User{Email: person.Email, Password: person.Password}
	fmt.Println(person)
	err = user.GetOne()
	 if err != nil {
		 w.WriteHeader(500)
		 fmt.Fprintln(w, err)
		 return
	 }

	 if ok := models.CheckPassword(person.Password, user.Password); !ok {
		 w.WriteHeader(401)
		 fmt.Fprintln(w, "unauthorized")
		 return
	 }
	
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": user.Email,
        "Id": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := unsignedToken.SignedString(Secret)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Sorry there was an error")
		return 
	}
	
	var tokenBody Token = Token{Token: tokenString}
	/*data, err := json.Marshal(tokenBody)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, err)
	}
	fmt.Println(data)
    fmt.Fprintln(w, data)*/
	json.NewEncoder(w).Encode(tokenBody)
    	
}