package main

import (
	"path"

	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", IndexRoute).Methods("GET")
	router.HandleFunc("/view", ViewWebsite).Methods("GET")
	router.HandleFunc("/edit", EditTitle).Methods("GET", "POST")
	router.HandleFunc("/save", EditBody).Methods("GET", "POST")
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	log.Println("ListenAndServe http://localhost:8090")
	if err := http.ListenAndServe(":8090", router); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "404.html", nil)
}

func ViewWebsite(w http.ResponseWriter, r *http.Request) {
	website := LoadWebsite()
	render(w, "view.html", website)
}

func EditTitle(w http.ResponseWriter, r *http.Request) {
	current := LoadWebsite()

	switch r.Method {
	case "POST":
		new := &Website{
			Title: r.PostFormValue("title"),
			Body:  current.Body,
		}

		if !new.Validate() {
			render(w, "edit_title.html", new)
			return
		}

		if err := new.Save(); err != nil {
			renderInternalError(w, err)
			return
		}

		render(w, "view.html", new)
	default:
		render(w, "edit_title.html", current)
	}
}

func EditBody(w http.ResponseWriter, r *http.Request) {
	current := LoadWebsite()

	switch r.Method {
	case "POST":
		new := &Website{
			Title: current.Title,
			Body:  r.PostFormValue("body"),
		}

		if !new.Validate() {
			render(w, "edit_body.html", new)
			return
		}

		if err := new.Save(); err != nil {
			renderInternalError(w, err)
			return
		}

		render(w, "view.html", new)
	default:
		render(w, "edit_body.html", current)
	}
}

func render(w http.ResponseWriter, html string, data interface{}) {
	templates := template.Must(template.ParseFiles("templates/layout.html", path.Join("templates", html)))
	templates.ExecuteTemplate(w, "layout", data)
}

func renderInternalError(w http.ResponseWriter, err error) {
	templates := template.Must(template.ParseFiles("templates/layout.html", path.Join("templates", "500.html")))
	templates.ExecuteTemplate(w, "layout", map[string]string{
		"Error": err.Error(),
	})
}
