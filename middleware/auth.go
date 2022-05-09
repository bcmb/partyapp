package middleware

import (
	"golang.org/x/crypto/bcrypt"
	"gotraining.com/user"
	"log"
	"net/http"
)

func BasicAuthMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		usr, pass, ok := r.BasicAuth()
		if !ok || !checkUsernameAndPassword(usr, pass) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorised"))
			return
		}
		handler(w, r)
	}
}

func checkUsernameAndPassword(username, password string) bool {
	usr, err := user.GetUser(username)

	if err != nil {
		log.Printf("Filed to fetch user %s", err)
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password))
		if err != nil {
			return false
		} else {
			return true
		}
	}
	return false
}
