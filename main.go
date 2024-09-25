package main

import (
	"GoLearn/internal/logic"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("static/templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func pagePasswordLen(w http.ResponseWriter, r *http.Request) {
	lenPassword := mux.Vars(r)["len"]  // get password length by url
	answer := logic.GeneratedPassword(lenPassword)
	tmpl, err := template.ParseFiles("static/templates/password_len.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, answer)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/{len:[0-9]+}", pagePasswordLen)

	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	router.PathPrefix("/static/").Handler(staticHandler)
	
	http.Handle("/", router)
	http.ListenAndServe(":8800", nil)
}
