package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"os"
	"github.com/go-chi/chi"
	"io"
	"image/png"
	"github.com/nfnt/resize"
	"golang.org/x/image/bmp"
	"github.com/naipath/mundiclient"
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

func getLaserClient(r *http.Request) LaserClient {
	laserClientName := chi.URLParam(r, "laserClientName")
	for _, x := range laserClients {
		if x.Name == laserClientName {
			return x
		}
	}
	return LaserClient{}
}

func convertPngToBmp(reader io.Reader, fileName string) error {
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

func getLaserClients(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(laserClients)
}

func addLaserClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var data CreateLaserClientRequest
	decodeErr := decoder.Decode(&data)
	if decodeErr != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"error\": \"Could not decode json request\"}"))
		return
	}
	newClient, clientErr := mundiclient.New(data.Ip, data.Port)
	if clientErr != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"Could not connect to Mundi laser\"}"))
	} else {
		newLaserClient := LaserClient{newClient, data.Name, data.Ip, data.Port}
		laserClients = append(laserClients, newLaserClient)

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(newLaserClient)
	}
}

func deleteLaserClient(w http.ResponseWriter, r *http.Request) {
	client := getLaserClient(r)
	client.client.Close()

	newClients := make([]LaserClient, 0)
	for _, v := range laserClients {
		if v.Name != client.Name {
			newClients = append(newClients, v)
		}
	}
	laserClients = newClients
	w.WriteHeader(204)
}

func getLaserClientCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	laserClient := getLaserClient(r)
	counters, err := laserClient.client.GetCounters()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"Could not retrieve counters from mundiclient\"}"))
	} else {
		json.NewEncoder(w).Encode(counters)
	}
}

func resetLaserClientCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	laserClient := getLaserClient(r)
	err := laserClient.client.ResetCurrentCount()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"Could not retrieve counters from mundiclient\"}"))
		return
	}
	w.WriteHeader(204)
}

func getLaserClientStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	laserClient := getLaserClient(r)
	statusData, statusDataErr := laserClient.client.GetStatusData()
	statusMessage, statusMessageErr := laserClient.client.GetStatusMessage()

	if statusDataErr != nil || statusMessageErr != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"Could not retrieve status from mundiclient\"}"))
		return
	}
	json.NewEncoder(w).Encode(LaserStatusRequest{statusMessage, statusData})
}

func uploadLogoToLaser(w http.ResponseWriter, r *http.Request) {
	client := getLaserClient(r)
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("{\"error\": \"Could not retrieve formfield uploadfile\"}"))
		return
	}
	fmt.Fprintf(w, "%v", handler.Header)

	fileName := *path + handler.Filename[:len(handler.Filename)-4] + ".bmp"
	convertErr := convertPngToBmp(file, fileName)
	if convertErr != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"Could not convert file to bmp\"}"))
		return
	}

	logo, _ := os.Open(fileName)

	uploadErr := client.client.UploadLogo(logo)
	if uploadErr != nil {
		w.WriteHeader(500)
		w.Write([]byte("{\"error\": \"Could not upload file to mundiclient\"}"))
		return
	}
	file.Close()
}

func writeError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}