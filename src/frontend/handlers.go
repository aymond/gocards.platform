package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// HomeHandler comment
func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)

	if err != nil {
		generateHTML(w, "", "home", "header.private", "footer")
	} else {
		generateHTML(w, "", "home", "header.private", "footer")
	}
}

// NameHandler comment
func nameHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("templates/*.html")
	t.ExecuteTemplate(w, "home", "")
}

// HealthCheckHandler comment
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", http.StatusOK)
}

// GameNewHandler comment
func gameNewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", "<h1>New Game!</h1>")
}

// GameListHandler comment
func gameListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", "<h1>Game List!</h1>")
}

// LoginHandler Comment
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", http.StatusOK)
}

// LogoutHandler comment
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", "<h1>New Game!</h1>")
}

// SignupHandler comment
func signupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", "<h1>Game List!</h1>")
}
