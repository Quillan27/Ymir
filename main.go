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

// Page holds the un-changing elements on the webpage
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
	// TODO(karl): turn into variable and set via settings menu
	WorldWidth int = 750

	// WorldHeight is the default world height
	// TODO(karl): turn into variable and set via settings menu
	WorldHeight int = 600
)

var (
	// tmpl is the HTML template for generating the webpage
	tmpl template.Template

	// world is the default current world being displayed
	world World
)

// pageHandler handles the loading of the webpages structure
func pageHandler(w http.ResponseWriter, r *http.Request) {
	// initialize the webpage struct
	p := Page{
		Title: PageTitle,
		Name:  "World Name",
	}

	// create a new template from the html layout
	tmpl, err := template.ParseFiles(HTMLPath)
	if err != nil {
		fmt.Print("ERROR: HTML template not found\n")
	}

	// write the template to the webpage
	tmpl.Execute(w, p)
}

// TODO(karl): Can these handles be consolidated?
// Can I check a button id to make this cleaner?

// newWorldHandler handles requests from the 'New World' button
func newWorldHandler(w http.ResponseWriter, r *http.Request) {
	world = *newWorld(WorldWidth, WorldHeight)
	w.Write([]byte(getEncodedMap()))
}

// worldNameHandler handles a request for a new world name
func worldNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>" + world.Name + "</h1>"))
}

// terrainViewHandler handles request's from the 'Terrain' button
func terrainViewHandler(w http.ResponseWriter, r *http.Request) {
	world.drawMap(TerrainView)
	w.Write([]byte(getEncodedMap()))
}

// biomeViewHandler handles request's from the 'Biome' button
func biomeViewHandler(w http.ResponseWriter, r *http.Request) {
	world.drawMap(BiomeView)
	w.Write([]byte(getEncodedMap()))
}

// politicalViewHandler handles request's from the 'Political' button
func politicalViewHandler(w http.ResponseWriter, r *http.Request) {
	world.drawMap(PoliticalView)
	w.Write([]byte(getEncodedMap()))
}

// climateViewHandler handles request's from the 'Climate' button
func climateViewHandler(w http.ResponseWriter, r *http.Request) {
	world.drawMap(ClimateView)
	w.Write([]byte(getEncodedMap()))
}

// getEncodedMap encodes the current world's map into a base64 image for HTML
func getEncodedMap() string {
	var buffer bytes.Buffer
	png.Encode(&buffer, &world.Map)
	i := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return "<img src=\"data:image/png;base64," + i + "\">"
}

func main() {
	// creating a new routing solution and adding handlers
	r := mux.NewRouter()

	// serve static files used by the webpage
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	// add various handlers for buttons and labels
	r.HandleFunc("/", pageHandler)
	r.HandleFunc("/newWorld", newWorldHandler)
	r.HandleFunc("/worldName", worldNameHandler)
	r.HandleFunc("/terrainView", terrainViewHandler)
	r.HandleFunc("/biomeView", biomeViewHandler)
	r.HandleFunc("/politicalView", politicalViewHandler)
	r.HandleFunc("/climateView", climateViewHandler)

	// open the server, report errors if needed
	fmt.Printf("Listening on :8080...\n")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
