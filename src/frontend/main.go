package main

import (
	"log"
	"net/http"
)

const (
	port             = "8080"
	cookiePrefix     = "cards_"
	cookieSessionsID = cookiePrefix + "session-id"
)

// version
func version() string {
	return "0.0.1"
}

func main() {
	p("gocards", version(), "started.")
	//r := mux.NewRouter()
	r := http.NewServeMux()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/HealthCheck", healthCheckHandler)

	// User Handlers
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/authenticate", authenticationHandler)
	r.HandleFunc("/newsignup", newSignup)
	r.HandleFunc("/profile", profile)

	// Other handlers for experimenting.
	r.HandleFunc("/name/{name}", nameHandler)

	// Game Handlers
	r.HandleFunc("/game/new", gameNewHandler)
	r.HandleFunc("/game/list", gameListHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
