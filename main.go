package main

import (
	"fmt"
	"net/http"
)

type Note struct {
	Id   int
	Text string
}

var notes []Note

func main() {

	notes = append(notes, Note{Id: 1, Text: "Test note"})
	http.HandleFunc("/notes", handleNotes)
	http.HandleFunc("/notes/", handleNotesById)
	http.ListenAndServe(":8080", nil)

}

func handleNotes(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		for _, value := range notes {
			fmt.Fprintf(writer, "%d - %s", value.Id, value.Text)
		}
	}
}

func handleNotesById(writer http.ResponseWriter, request *http.Request) {

}
