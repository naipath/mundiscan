package main

import (
	"flag"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

var (
	path *string
	settingsFile *string
)

const noSettings = "no-settings"

func main() {
	port := flag.String("port", "8100", "port to serve on")
	directory := flag.String("statics", ".", "the directory of static files to host")
	path = flag.String("path", "./", "The path to write file uploads to")
	settingsFile = flag.String("settings", noSettings, "The path to read and write the settings file to, which contains the laserClients")
	flag.Parse()

	InitializeLaserClients(*settingsFile)

	r := chi.NewRouter()

	r.Route("/laserclients", func(r chi.Router) {
		r.Get("/", getLaserClients)
		r.Post("/", addLaserClient)

		r.Route("/{laserClientId}", func(r chi.Router) {
			r.Get("/", retrieveLaserClient)
			r.Delete("/", deleteLaserClient)
			r.Get("/status", getLaserClientStatus)
			r.Post("/upload", uploadLogoToLaser)

			r.Route("/count", func(r chi.Router) {
				r.Get("/", getLaserClientCount)
				r.Delete("/", resetLaserClientCount)
			})
		})
	})

	r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(*directory)).ServeHTTP(w, r)
	}))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":" + *port, r))
}