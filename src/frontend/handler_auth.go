package main

import (
	"net/http"
	"time"

	"./data"
)

// Session comment
type Session struct {
	ID    int
	UUID  string
	Email string

	UserID    int
	CreatedAt time.Time
}

func authenticationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, _ := user.CreateSession()
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

// GET /login
// Show the login page
func login(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("footer", "header.public", "login")
	t.Execute(writer, nil)
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
	t := parseTemplateFiles("footer", "header.public", "signup")
	t.Execute(writer, nil)
	//generateHTML(writer, nil, "footer", "header.public", "")
}

// POST /signup
// newSignup Create the user account
func newSignup(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		logerror(err, "Cannot parse form")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		logerror(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}
