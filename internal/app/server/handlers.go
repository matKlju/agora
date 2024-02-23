package server

import (
	"Forum/internal/model"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("./web/templates/*.html"))

func (s *server) HandlePaths() {
	s.router.Handle("/static/", s.serveStatic())
	s.router.HandleFunc("/", s.home())
	s.router.HandleFunc("/registerPage", s.registerPage())
	s.router.HandleFunc("/saveUser", s.saveRegister())
	s.router.HandleFunc("/login", s.login())
}

func (s *server) registerPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		execTmpl(w, templates.Lookup("register.html"), nil)
	}
}

func (s *server) saveRegister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Println("@ register page")

		userName := r.FormValue("userName")
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := model.NewUser(userName, email, password)
		if err != nil {
			http.Redirect(w, r, "/registerPage", http.StatusBadRequest)
			s.logger.Println("redirect - NewUser() error: ", err)
			return
		}
		/*
			existingUser, err := s.store.User().GetByUsername(userName)
			if err != nil && err != sql.ErrNoRows {
				http.Error(w, "Error checking for existing user", http.StatusInternalServerError)
				s.logger.Println("GetByUsername() error: ", err)
				return
			}

			if existingUser != nil {
				// Handle the case where the username already exists
				http.Error(w, "Username already taken", http.StatusBadRequest)
				return
			}

			// Continue with user registration...
		*/

		if err = s.store.User().Register(user); err != nil {
			s.logger.Println("redirect - Register() error: ", err)
			http.Redirect(w, r, "/registerPage", http.StatusBadRequest)
			return
		}

		execTmpl(w, templates.Lookup("main.html"), nil)
	}
}

func (s *server) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Println("@ login page")
		execTmpl(w, templates.Lookup("login.html"), nil)
	}
}

func (s *server) serveStatic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))).ServeHTTP(w, r)
	}
}

func (s *server) home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Println(" @ Homepage")
		execTmpl(w, templates.Lookup("main.html"), nil)
	}
}

// execTmpl renders a template with the given data or returns an internal server error.
func execTmpl(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing template:", err)
	}
}
