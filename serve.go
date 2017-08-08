package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/naipath/mundiclient"
	"encoding/json"
	"image/png"
	"golang.org/x/image/bmp"
	"io"
	"github.com/nfnt/resize"
	"github.com/go-chi/chi"
)

type LaserClient struct {
	client *mundiclient.MundiClient
	Name   string
	Ip     string
	Port   int
}

var laserClients = make([]LaserClient, 0)

type CreateLaserClientRequest struct {
	Name string
	Ip   string
	Port int
}

type LaserStatusRequest struct {
	Message string
	Status  mundiclient.StatusData
}

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	path := flag.String("path", "./", "The path to write file uploads to")
	flag.Parse()

	r := chi.NewRouter()

	r.Route("/laserclients", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(laserClients)
		})

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(laserClients)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			decoder := json.NewDecoder(r.Body)
			var data CreateLaserClientRequest
			decodeErr := decoder.Decode(&data)
			if decodeErr != nil {
				w.WriteHeader(400)
				w.Write([]byte("{\"error\": \"Could not decode json request\"}"))
			}
			newClient, clientErr := mundiclient.New(data.Ip, data.Port)
			if clientErr != nil {
				w.WriteHeader(500)
				w.Write([]byte("{\"error\": \"Could not connect to Mundi laser\"}"))
			}

			newLaserClient := LaserClient{newClient, data.Name, data.Ip, data.Port}
			laserClients = append(laserClients, newLaserClient)

			w.WriteHeader(201)
			json.NewEncoder(w).Encode(newLaserClient)
		})

		r.Route("/{laserClientName}", func(r chi.Router) {
			r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
				b := laserClients[:0]
				b = append(b, getLaserClient(r))
				laserClients = b

				w.WriteHeader(204)
			})

			r.Get("/count", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")

				client := getLaserClient(r)
				counters, err := client.client.GetCounters()
				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte("{\"error\": \"Could not retrieve counters from mundiclient\"}"))
				} else {
					json.NewEncoder(w).Encode(counters)
				}
			})

			r.Route("/count", func(r chi.Router) {
				r.Get("/", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")

					laserClient := getLaserClient(r)
					counters, err := laserClient.client.GetCounters()
					if err != nil {
						w.WriteHeader(500)
						w.Write([]byte("{\"error\": \"Could not retrieve counters from mundiclient\"}"))
					} else {
						json.NewEncoder(w).Encode(counters)
					}
				})

				r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")

					laserClient := getLaserClient(r)
					err := laserClient.client.ResetCurrentCount()
					if err != nil {
						w.WriteHeader(500)
						w.Write([]byte("{\"error\": \"Could not retrieve counters from mundiclient\"}"))
					}
					w.WriteHeader(204)
				})
			})

			r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")

				laserClient := getLaserClient(r)
				statusData, statusDataErr := laserClient.client.GetStatusData()
				statusMessage, statusMessageErr := laserClient.client.GetStatusMessage()

				if statusDataErr != nil || statusMessageErr != nil {
					w.WriteHeader(500)
					w.Write([]byte("{\"error\": \"Could not retrieve status from mundiclient\"}"))
				}
				json.NewEncoder(w).Encode(LaserStatusRequest{statusMessage, statusData})
			})

			r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
				client := getLaserClient(r)
				r.ParseMultipartForm(32 << 20)
				file, handler, err := r.FormFile("uploadfile")
				if err != nil {
					w.WriteHeader(400)
					w.Write([]byte("{\"error\": \"Could not retrieve formfield uploadfile\"}"))
					fmt.Println(err)
					return
				}
				fmt.Fprintf(w, "%v", handler.Header)

				fileName := *path + handler.Filename[:len(handler.Filename)-4] + ".bmp"
				convertErr := convertPngToBmp(file, fileName)
				if convertErr != nil {
					w.WriteHeader(500)
					w.Write([]byte("{\"error\": \"Could not convert file to bmp\"}"))
					fmt.Println(convertErr)
					return
				}

				logo, _ := os.Open(fileName)

				uploadErr := client.client.UploadLogo(logo)
				if uploadErr != nil {
					w.WriteHeader(500)
					w.Write([]byte("{\"error\": \"Could not upload file to mundiclient\"}"))
					fmt.Println(uploadErr)
					return
				}
				file.Close()
			})

		})
	})

	r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(*directory)).ServeHTTP(w, r)
	}))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":" + *port, r))
}

func getLaserClient(r *http.Request) LaserClient {
	laserClientName := chi.URLParam(r, "laserClientName")
	for _, x := range laserClients {
		if x.Name == laserClientName {
			return x
		}
	}
	return LaserClient{}
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