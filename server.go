package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	//	"html/template"
	"image/png"
	"log"
	"net/http"
)

func mapHandler(w http.ResponseWriter, r *http.Request) {
	world := newWorld(512, 512)

	var buffer bytes.Buffer
	png.Encode(&buffer, &world.Map)
	encodedImage := base64.StdEncoding.EncodeToString(buffer.Bytes())
	html := "<html><body><img src=\"data:image/png;base64," + encodedImage + "\" /></body></html>"

	w.Write([]byte(fmt.Sprintf(html)))
}

func startServer() {
	http.HandleFunc("/", mapHandler)

	fmt.Printf("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
