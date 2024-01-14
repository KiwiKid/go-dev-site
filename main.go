package main

import (
	"context"
	"log"
	"os"
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

	outputDir := "./output"

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755) // 0755 is a common permission setting (read and execute access for everyone, full access for the owner)
		if err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
	}

	f, err := os.Create(outputDir + "/home.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = home().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}

	/*component :=

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
	}*/
}
