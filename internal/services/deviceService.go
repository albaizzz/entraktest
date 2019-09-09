package services

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/beevik/etree"
	"github.com/entraktest/internal/models"
	"github.com/entraktest/internal/repositories"
	"github.com/entraktest/pkg/constants"
)

type DeviceService struct {
	DeviceRepository repositories.IDeviceRepository
}

type IDeviceService interface {
	Store(dataXml []byte) error
	GetById(id uint64) (models.Device, error)
	GetAll() ([]models.Device, error)
}

func NewDeviceService(deviceRepository repositories.IDeviceRepository) *DeviceService {
	return &DeviceService{
		DeviceRepository: deviceRepository,
	}
}

func (ds *DeviceService) GetById(id uint64) (device models.Device, err error) {
	device, err = ds.DeviceRepository.GetById(id)
	return
}

func (ds *DeviceService) GetAll() (devices []models.Device, err error) {
	devices, err = ds.DeviceRepository.GetAll()
	return
}

func (ds *DeviceService) Store(dataXml []byte) error {
	var deviceRequest models.DeviceRequest
	xml.Unmarshal(dataXml, &deviceRequest)
	doc := etree.NewDocument()
	if err := doc.ReadFromString(string(dataXml)); err != nil {
		panic(err)
	}
	root := doc.SelectElement("root")
	devices := root.SelectElement("devices")

	for _, deviceMod := range deviceRequest.Devicemodel {
		var device models.Device
		device.Device = deviceMod.Device
		device.Value = deviceMod.Value
		device.Datetime = time.Unix(deviceRequest.RecordTime, 0)

		deviceID := devices.SelectElement(deviceMod.Device)
		if deviceID == nil {
			return fmt.Errorf("Unmatch data request")
		}
		switch deviceID.Text() {
		case constants.PowerMeter:
			device.Unit = "A"
		case constants.VoltageMeter:
			device.Unit = "V"
		case constants.CurrentMeter:
			device.Unit = "C"
		case constants.TemperatureSensor:
			device.Unit = "C"
		}
		err := ds.DeviceRepository.Store(device)
		if err != nil {
			return fmt.Errorf("Error when inserting devices, %s", err)
		}
	}
	return nil
}
