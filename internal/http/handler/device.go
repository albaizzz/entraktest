package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/entraktest/internal/services"
	"github.com/entraktest/pkg/responses"
	"github.com/gorilla/mux"
)

type DeviceHandler struct {
	DeviceService services.IDeviceService
}

type IDeviceHandler interface {
	InsertDevice(w http.ResponseWriter, r *http.Request)
	GetDevice(w http.ResponseWriter, r *http.Request)
	GetDevices(w http.ResponseWriter, r *http.Request)
}

func NewDeviceHandler(deviceService services.IDeviceService) *DeviceHandler {
	return &DeviceHandler{
		DeviceService: deviceService,
	}
}

func (d *DeviceHandler) InsertDevice(w http.ResponseWriter, r *http.Request) {

	dataXml, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	err = d.DeviceService.Store(dataXml)
	if err != nil {
		responses.Write(w, responses.APIErrorUnknown)
	}
	responses.Write(w, responses.APICreated)
}

func (d *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, _ := strconv.ParseUint(param["id"], 10, 64)
	device, err := d.DeviceService.GetById(id)
	if err != nil {
		responses.Write(w, responses.APIErrorUnknown)
	}
	responses.Write(w, responses.APIOK.WithData(device))
}

func (d *DeviceHandler) GetDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := d.DeviceService.GetAll()
	if err != nil {
		responses.Write(w, responses.APIErrorUnknown)
	}
	responses.Write(w, responses.APIOK.WithData(devices))
}
