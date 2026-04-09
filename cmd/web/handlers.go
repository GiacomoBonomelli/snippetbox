package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
// Use the template.ParseFiles() function to read the template file into a
// template set. If there's an error, we log the detailed error message, use
// the http.Error() function to send an Internal Server Error response to the
// user, and then return from the handler so no subsequent code is executed.
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
	return
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// WriteHeader va prima di Write perché quest'ultima invia automaticamente lo status code 200.
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Save a new snippet"))
}
