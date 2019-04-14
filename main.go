package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"image/png"
	"log"
	"net/http"
)

// aspects of the base webpage, loaded with an HTML template
type Page struct {
	Title string // webpage title
	Name  string // world name
}

const (
	HTML_PATH  string = "static/layout.html" // location HTML template
	PAGE_TITLE string = "Ymir"               // webpage title
)

var (
	tmpl  template.Template // HTML template for the webpage
	world World             // the currently loaded world
)

// handles loading the base webpage with an HTML template
func pageHandler(w http.ResponseWriter, r *http.Request) {
	// initialize the webpage struct
	//
	page := Page{
		Title: PAGE_TITLE,
		Name:  "World Name",
	}

	tmpl, err := template.ParseFiles(HTML_PATH)
	if err != nil {
		fmt.Print("jemplate not found\n")
	}

	tmpl.Execute(w, page)
}

// creates a totally new world and passes it's map back to the webpage
// linked to the new world button on the webpage
func newWorldHandler(w http.ResponseWriter, r *http.Request) {
	// create a totally new world
	world = *newWorld(512, 512)

	// encode the image.RGBA to a base64 encoded image
	var buffer bytes.Buffer
	png.Encode(&buffer, &world.Map)
	encodedImage := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// pass the encoded image back enclosed in HTML
	w.Write([]byte("<img src=\"data:image/png;base64," + encodedImage + "\">"))
}

func main() {
	// creating a new routing solution and add handlers
	router := mux.NewRouter()
	router.HandleFunc("/", pageHandler)
	router.HandleFunc("/map", newWorldHandler)

	// open the server, report errors if needed
	fmt.Printf("Listening on :8080...\n")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
