package bluetooth

import (
	// std
	"time"

	// external
	"github.com/google/uuid"
	"tinygo.org/x/bluetooth"
)


func ConvertToBluetoothAdvertisement(bleAdvertisement *BluetoothAdvertisement) (*bluetooth.Advertisement, error) {
	var serviceUUIDs []bluetooth.UUID

	if bleAdvertisement.AdvertiseData.ServiceUuid != uuid.Nil {
		uuid, err := bleAdvertisement.AdvertiseData.ServiceUuid.MarshalBinary()
		if err != nil {
			return nil, err
		}

		var uuidSliced[16]byte
		for i := 0; i < len(uuidSliced); i++ {
			uuidSliced[i] = uuid[i]
		}

		serviceUUIDs = append(serviceUUIDs, bluetooth.NewUUID(uuidSliced))
	}

	if bleAdvertisement.AdvertiseData.ServiceSolicitationUuid != uuid.Nil {
		uuid, err := bleAdvertisement.AdvertiseData.ServiceSolicitationUuid.MarshalBinary()
		if err != nil {
			return nil, err
		}

		var uuidSliced[16]byte
		for i := 0; i < len(uuidSliced); i++ {
			uuidSliced[i] = uuid[i]
		}

		serviceUUIDs = append(serviceUUIDs, bluetooth.NewUUID(uuidSliced))
	}

	time, err := time.ParseDuration(bleAdvertisement.AdvertiseSettings.Interval)
	if (err != nil) {
		return nil, err
	}

	advertise := bluetooth.DefaultAdapter.DefaultAdvertisement()
	advertise.Configure(
		bluetooth.AdvertisementOptions{
				LocalName: bleAdvertisement.AdvertiseSettings.Name,
				Interval: bluetooth.NewDuration(time),
				ManufacturerData: convertToManufacturerDataElement(bleAdvertisement.AdvertiseData.ManufacturerData),
				ServiceUUIDs: serviceUUIDs,
				ServiceData: convertToServiceDataElement(bleAdvertisement.AdvertiseData.ServiceData),
			},
	)

	return advertise, nil
}

func convertToManufacturerDataElement(manufacturerData []*ManufacturerDataModel) []bluetooth.ManufacturerDataElement {
	size := len(manufacturerData)
	var manufacturersElements []bluetooth.ManufacturerDataElement

	for i := 0; i < size; i++ {
		manufacturersElements = append(manufacturersElements, bluetooth.ManufacturerDataElement{
			CompanyID: manufacturerData[i].Id,
			Data: manufacturerData[i].Data,
		})
	}

	return manufacturersElements
}

func convertToServiceDataElement(serviceData []*ServiceDataModel) []bluetooth.ServiceDataElement {
	size := len(serviceData)
	var servicesElements []bluetooth.ServiceDataElement

	for i := 0; i < size; i++ {
		servicesElements = append(servicesElements, bluetooth.ServiceDataElement{
			UUID: convertUuid(serviceData[i].Uuid),
			Data: serviceData[i].Data,
		})
	}

	return servicesElements
}

func convertUuid(uuid uuid.UUID) bluetooth.UUID {
	uuidDecoded, err := uuid.MarshalBinary()
	if (err != nil) {
		panic("Failed to convert UUID" + err.Error())
	}
	return bluetooth.NewUUID([16]byte(uuidDecoded))
}