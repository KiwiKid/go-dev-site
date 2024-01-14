package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

/*
func setMIMEType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "" {
			switch {
			case r.URL.Path[len(r.URL.Path)-4:] == ".css":
				w.Header().Set("Content-Type", "text/css")
				// You can add other file type detections here if needed
			}
		}
		next.ServeHTTP(w, r)
	})
}*/

func main() {

	component := home()

	http.Handle("/", templ.Handler(component))

	aboutPage := aboutWithContainer()

	http.Handle("/about", templ.Handler(aboutPage))

	debug := debug()

	http.Handle("/debug", templ.Handler(debug))

	nzconfigmap := nzCovidMapWithContainer()

	http.Handle("/dev/nzcovidmap", templ.Handler(nzconfigmap))

	assetsHandler := http.StripPrefix("/assets/", http.FileServer(http.Dir("assets")))

	// apply the middleware to the assets handler
	http.Handle("/assets/", assetsHandler)

	fmt.Println("Listening on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Printf("error listening: %v", err)
	}
}
