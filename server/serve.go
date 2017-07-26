package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/naipath/mundiclient"
	"encoding/json"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	path := flag.String("path", "./test/", "The path to write file uploads to")
	flag.Parse()

	client, err := mundiclient.New("10.195.0.171", 1470)
	if err != nil {
		panic("Could not get a connection to Mundi laser")
	}
	defer client.Close()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("method:", r.Method)
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(*path + handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	})

	http.HandleFunc("/statusData", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		statusData, err := client.GetStatusData()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("{\"error\": true}"))
		} else {
			json.NewEncoder(w).Encode(statusData)
		}
	})

	http.HandleFunc("/statusMessage", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		statusMessage, err := client.GetStatusMessage()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("{\"error\": true}"))
		} else {
			w.Write([]byte("{\"StatusMessage\": \"" + statusMessage + "\"}"))
		}
	})

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}
