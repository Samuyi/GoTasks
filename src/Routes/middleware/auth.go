package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)


type User struct {
	Email string 
	Id int
}

var Secret = []byte("samuyi902210")

func Auth(f http.HandlerFunc) http.HandlerFunc {
	  return func (w http.ResponseWriter, r *http.Request) {
		  bearerToken := r.Header.Get("Authorization")
		  if bearerToken == "" {
			   w.WriteHeader(401)
			   fmt.Fprintln(w, "unauthorized")
			   return
		  }
		  tokenString := bearerToken[7:]
		  token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
              if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error") 
			  }
			  return Secret, nil
		  })
          if err != nil {
			 w.WriteHeader(401)
			 return
		  }

		  if token.Valid {
		  var user User
		  mapstructure.Decode(token.Claims, &user)
		  var id string = strconv.Itoa(user.Id) 
		  r.Header.Set("email", user.Email)
		  r.Header.Set("id", id)
          f(w, r)
		  }
	  }
  }

