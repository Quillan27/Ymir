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

type Page struct {
	Title string
	Name  string
}

const (
	HTML_PATH  string = "static/layout.html"
	PAGE_TITLE string = "Ymir"
)

var (
	tmpl  template.Template
	world World
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title: PAGE_TITLE,
		Name:  "World Name",
	}

	tmpl, err := template.ParseFiles(HTML_PATH)
	if err != nil {
		fmt.Print("Template not found\n")
	}

	tmpl.Execute(w, page)
}

func newWorldHandler(w http.ResponseWriter, r *http.Request) {
	world = *newWorld(512, 512)

	var buffer bytes.Buffer
	png.Encode(&buffer, &world.Map)
	encodedImage := base64.StdEncoding.EncodeToString(buffer.Bytes())

	w.Write([]byte("<img src=\"data:image/png;base64," + encodedImage + "\">"))
}

func startServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", pageHandler)
	router.HandleFunc("/map", newWorldHandler)

	fmt.Printf("Listening on :8080\n")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
