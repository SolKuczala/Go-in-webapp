package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/*port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}*/
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/sing-up", signUp)              //post
	router.HandleFunc("/api/v1/sing-up/google", signUpGoogle) //post
	router.HandleFunc("/api/v1/login", login)                 //
	router.HandleFunc("/api/v1/login/google", loginGoogle)
	router.HandleFunc("/api/v1/user", userGetProfile)      //search get mux
	router.HandleFunc("/api/v1/user", userEditProfile)     //search post mux
	router.HandleFunc("/api/v1/logout", logout)            //post
	router.HandleFunc("/api/v1/reset-pass", resetPassword) //must handle forgot/reset

	http.Handle("/", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func signUp() {

}

func signUpGoogle() {

}

func login() {

}

func loginGoogle() {

}

func userGetProfile() {

}

func userGetProfile() {

}

func userEditProfile() {

}

func logout() {

}

func resetPassword() {

}
