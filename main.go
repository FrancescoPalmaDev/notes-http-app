package main

import "net/http"

func main() {
	http.HandleFunc("/notes", handleNotes)
	http.HandleFunc("/notes/", handleNotesById)
	http.ListenAndServe(":8080", nil)
}

func handleNotes(writer http.ResponseWriter, request *http.Request) {

}

func handleNotesById(write http.ResponseWriter, request *http.Request) {

}
