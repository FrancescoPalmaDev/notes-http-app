package main

import (
	"fmt"
	"io"
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
			fmt.Fprintf(writer, "%d - %s\n", value.Id, value.Text)
		}
	case "POST":
		body, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Fprint(writer, "Not able to read the body")
			return
		}
		newNote := Note{Id: len(notes) + 1, Text: string(body)}
		notes = append(notes, newNote)
		fmt.Fprintf(writer, "Created note N-%d\n", newNote.Id)

	}
}

func handleNotesById(writer http.ResponseWriter, request *http.Request) {

}
