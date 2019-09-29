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

func main() {
	//r := mux.NewRouter()
	r := http.NewServeMux()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/HealthCheck", healthCheckHandler)

	// User Handlers
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/signup", signupHandler)
	r.HandleFunc("/authenticate", authenticationHandler)

	// Other handlers for experimenting.
	r.HandleFunc("/name/{name}", nameHandler)

	// Game Handlers
	r.HandleFunc("/game/new", gameNewHandler)
	r.HandleFunc("/game/list", gameListHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
