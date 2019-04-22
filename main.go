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

// Page holds the elements on the webpage
type Page struct {
	Title string // webpage title
	Name  string // world name
}

const (
	HTML_PATH    string = "static/layout.html" // location of HTML template
	PAGE_TITLE   string = "Ymir"               // webpage title
	WORLD_WIDTH  int    = 512
	WORLD_HEIGHT int    = 512
)

var (
	tmpl  template.Template // HTML template for the webpage
	world World             // the currently loaded world
)

// pageHandler handles the loading of the webpages structure
func pageHandler(w http.ResponseWriter, r *http.Request) {
	// initialize the webpage struct
	page := Page{
		Title: PAGE_TITLE,
		Name:  "World Name",
	}

	// create a new template from the html layout
	tmpl, err := template.ParseFiles(HTML_PATH)
	if err != nil {
		fmt.Print("jemplate not found\n")
	}

	// write the template to the webpage
	tmpl.Execute(w, page)
}

// newWorldHandler handles requests from the 'New World' button
// it generates a totally new world and passes back a map
func newWorldHandler(w http.ResponseWriter, r *http.Request) {
	// create a totally new world
	world = *newWorld(WORLD_WIDTH, WORLD_HEIGHT)

	// encode the image.RGBA to a base64 encoded image
	var buffer bytes.Buffer
	png.Encode(&buffer, &world.Map)
	encodedImage := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// pass the encoded image back enclosed in HTML
	w.Write([]byte("<img src=\"data:image/png;base64," + encodedImage + "\">"))
}

func main() {
	// creating a new routing solution and adding handlers
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
