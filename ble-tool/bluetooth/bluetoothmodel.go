package bluetooth


import (
	// external
	"github.com/google/uuid"
)


type ManufacturerDataModel struct {
	Id uint16
	Data []byte
}

type ServiceDataModel struct {
	Uuid uuid.UUID
	Data []byte
}

type BleAdvertisingSettings struct {
	Name string
	IncludeTxPower bool
	IncludeDeviceName bool
	Interval string
	Connectable bool
	LegacyMode bool
	Scannable bool
}

func NewBleAdvertisingSettings(
	name string,
	includeTxPower *bool,
	includeDeviceName *bool,
	interval *string,
	connectable *bool,
	legacyMode *bool,
	scannable *bool,
) *BleAdvertisingSettings {
	var bleAdvertisingSettings = &BleAdvertisingSettings{
		Name: name,
		IncludeTxPower: true,
		IncludeDeviceName: true,
		Interval: "250ms",
		Connectable: false,
		LegacyMode: true,
		Scannable: true,
	}

	if includeTxPower != nil {
		bleAdvertisingSettings.IncludeTxPower = *includeTxPower
	}

	if includeDeviceName != nil {
		bleAdvertisingSettings.IncludeDeviceName = *includeDeviceName
	}

	if interval != nil {
		bleAdvertisingSettings.Interval = *interval
	}

	if connectable != nil {
		bleAdvertisingSettings.Connectable = *connectable
	}

	if scannable != nil {
		bleAdvertisingSettings.Scannable = *scannable
	}

	return bleAdvertisingSettings
}

type BleAdvertisingData struct {
	ServiceUuid uuid.UUID
	ServiceSolicitationUuid uuid.UUID
	ManufacturerData []*ManufacturerDataModel
	ServiceData []*ServiceDataModel
}

func NewBleAdvertisingData(
	serviceUuid *uuid.UUID,
	serviceSolicitationUuid *uuid.UUID,
	manufacturerData []*ManufacturerDataModel,
	serviceData []*ServiceDataModel,
) *BleAdvertisingData {
	var bleAdvertisingData = &BleAdvertisingData{}

	if serviceUuid != nil {
		bleAdvertisingData.ServiceUuid = *serviceUuid
	}

	if serviceSolicitationUuid != nil {
		bleAdvertisingData.ServiceSolicitationUuid = *serviceSolicitationUuid
	}

	if manufacturerData != nil {
		bleAdvertisingData.ManufacturerData = manufacturerData
	}

	if serviceData != nil {
		bleAdvertisingData.ServiceData = serviceData
	}

	return bleAdvertisingData
}

type BluetoothAdvertisement struct {
	AdvertiseSettings *BleAdvertisingSettings
	AdvertiseData *BleAdvertisingData
}