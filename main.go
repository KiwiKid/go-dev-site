package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {



    component := home()
    
    http.Handle("/", templ.Handler(component))

	aboutPage := aboutPage()

	http.Handle("/about", templ.Handler(aboutPage))

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))


    fmt.Println("Listening on :3000")
    if err := http.ListenAndServe(":3000", nil); err != nil {
        log.Printf("error listening: %v", err)
    }
}