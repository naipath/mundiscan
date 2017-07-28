package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/naipath/mundiclient"
	"encoding/json"
	"strconv"
	"image/png"
	"golang.org/x/image/bmp"
	"io"
	"github.com/nfnt/resize"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	path := flag.String("path", "./", "The path to write file uploads to")
	flag.Parse()

	clientPort := 1470
	clientIp := "10.195.0.171"
	client, err := mundiclient.New(clientIp, clientPort)
	if err != nil {
		panic("Could not get a connection to Mundi laser")
	}
	client.SetDebug(true)
	defer client.Close()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)


		fileName := *path + handler.Filename[:len(handler.Filename) - 4] + ".bmp"
		convertErr := convertPngToBmp(file, fileName)
		if convertErr != nil {
			fmt.Println(convertErr)
			return
		}

		logo, _ := os.Open(fileName)

		uploadErr := client.UploadLogo(logo)
		if uploadErr != nil {
			fmt.Println(uploadErr)
			return
		}
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

	http.HandleFunc("/counters", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		counters, err := client.GetCounters()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("{\"error\": true}"))
		} else {
			json.NewEncoder(w).Encode(counters)
		}
	})

	http.HandleFunc("/resetCurrentCount", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := client.ResetCurrentCount()
		if err != nil {
			w.WriteHeader(500)
		}
	})

	http.HandleFunc("/clientSettings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.Write([]byte("{\"port\": \"" + strconv.Itoa(clientPort) + "\", \"ip\": \"" + clientIp + "\"}"))
	})

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}

func convertPngToBmp(reader io.Reader, fileName string) (error) {
	img, decodeErr := png.Decode(reader)

	resizedImg := resize.Resize(210, 210, img, resize.Lanczos3)

	if decodeErr != nil {
		return decodeErr
	}
	f, createErr := os.Create(fileName)
	defer f.Close()
	if createErr != nil {
		return createErr
	}
	bmp.Encode(f, resizedImg)
	return nil
}

//
//func convertPngToBmp(reader io.Reader, fileName string) (error) {
//	img, decodeErr := png.Decode(reader)
//
//	//resizedImg := resize.Resize(210, 210, img, resize.Lanczos3)
//
//	if decodeErr != nil {
//		return decodeErr
//	}
//	f, createErr := os.Create(fileName)
//	defer f.Close()
//	if createErr != nil {
//		return createErr
//	}
//	bmp.Encode(f, img)
//	return nil
//}