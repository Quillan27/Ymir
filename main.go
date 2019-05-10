package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image/png"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Page holds the elements on the webpage
type Page struct {
	Title string
	Name  string
}

const (
	// HTMLPath holds the location of HTML template
	HTMLPath string = "static/layout.html"

	// PageTitle is the web-browsers title
	PageTitle string = "Ymir"

	// WorldWidth is the default world width
	WorldWidth int = 512

	// WorldHeight is the default world height
	WorldHeight int = 512
)

var (
	tmpl  template.Template
	world World
)

// pageHandler handles the loading of the webpages structure
func pageHandler(w http.ResponseWriter, r *http.Request) {
	// initialize the webpage struct
	page := Page{
		Title: PageTitle,
		Name:  "World Name",
	}

	// create a new template from the html layout
	tmpl, err := template.ParseFiles(HTMLPath)
	if err != nil {
		fmt.Print("ERROR: HTML template not found\n")
	}

	// write the template to the webpage
	tmpl.Execute(w, page)
}

// handles requests from the 'New World' button
func newWorldHandler(w http.ResponseWriter, r *http.Request) {
	world = *newWorld(WorldWidth, WorldHeight)

	var buffer bytes.Buffer
	png.Encode(&buffer, &world.Map)
	encodedImage := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// write the encoded image to HTML
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
