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

type Page struct {
	Title string
	Name  string
}

const (
	HTMLPath    string = "static/layout.html"
	PageTitle   string = "Ymir"
	WorldWidth  int    = 1000
	WorldHeight int    = 800
)

var (
	tmpl template.Template

	world World
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title: PageTitle,
		Name:  "-",
	}

	tmpl, err := template.ParseFiles(HTMLPath)
	if err != nil {
		fmt.Print("ERROR: HTML template not found\n")
	}

	tmpl.Execute(w, p)
}

func newWorldHandler(w http.ResponseWriter, r *http.Request) {
	world = *newWorld(WorldWidth, WorldHeight)
	w.Write([]byte(getEncodedMap()))
}

func worldNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>" + world.Name + "</h1>"))
}

func elevationViewHandler(w http.ResponseWriter, r *http.Request) {
	world.drawMap(ElevationView)
	w.Write([]byte(getEncodedMap()))
}

func topographyViewHandler(w http.ResponseWriter, r *http.Request) {
	world.drawMap(TopographyView)
	w.Write([]byte(getEncodedMap()))
}

func biomeViewHandler(w http.ResponseWriter, r *http.Request) {
	world.drawMap(BiomeView)
	w.Write([]byte(getEncodedMap()))
}

func climateViewHandler(w http.ResponseWriter, r *http.Request) {
	world.drawMap(ClimateView)
	w.Write([]byte(getEncodedMap()))
}

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
	r.HandleFunc("/elevationView", elevationViewHandler)
	r.HandleFunc("/topographyView", topographyViewHandler)
	r.HandleFunc("/biomeView", biomeViewHandler)
	r.HandleFunc("/climateView", climateViewHandler)

	// open the server, report errors if needed
	fmt.Printf("Listening on :8080...\n")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
