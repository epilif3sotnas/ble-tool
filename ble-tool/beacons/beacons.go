package beacons


import (
	// std
	"encoding/binary"

	// internal
	"ble-tool/bluetooth"

	// external
	"github.com/google/uuid"
)


type iBeacon struct {
	Uuid uuid.UUID
	Major int16
	Minor int16
}

func NewiBeacon(uuid uuid.UUID, major int16, minor int16) *iBeacon {
	return &iBeacon {
		Uuid: uuid,
		Major: major,
		Minor: minor,
	}
}

func (self *iBeacon) ConvertToBluetoothAdvertisement(name string) *bluetooth.BluetoothAdvertisement {
	var bluetoothAdvertisement = &bluetooth.BluetoothAdvertisement{
		AdvertiseSettings: bluetooth.NewBleAdvertisingSettings(
			name,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		),
	}

	var manufacturerData = []byte {
		0x02, 0x15,
	}

	uuid, err := self.Uuid.MarshalBinary()
	if err != nil {
		panic("Failed to convert UUID" + err.Error())
	}
	manufacturerData = append(manufacturerData, uuid...)
	manufacturerData = binary.BigEndian.AppendUint16(manufacturerData, uint16(self.Major))
	manufacturerData = binary.BigEndian.AppendUint16(manufacturerData, uint16(self.Minor))
	manufacturerData = append(manufacturerData, 0xf9)

	bluetoothAdvertisement.AdvertiseData = &bluetooth.BleAdvertisingData{
		ManufacturerData: []*bluetooth.ManufacturerDataModel{
			{
				Id: 0x004c,
				Data: manufacturerData,
			},
		},
	}

	return bluetoothAdvertisement
}


type EddystoneFrame interface {
	Bytes() []byte
}

type EddystoneUID struct {
	NamespaceId [10]byte
	InstanceId [6]byte
	Rfu [2]byte
}

func NewEddystoneUID(
	namespaceId [10]byte,
	instanceId [6]byte,
	rfu [2]byte,
) *EddystoneUID {
	return &EddystoneUID{
		NamespaceId: namespaceId,
		InstanceId: instanceId,
		Rfu: rfu, }
}

func (self *EddystoneUID) Bytes() []byte {
	var frameType uint8 = 0x00
	var power uint8 = 0xf9

	var data = []byte{}
	data = append(data, frameType)
	data = append(data, power)
	data = append(data, self.NamespaceId[:]...)
	data = append(data, self.InstanceId[:]...)
	data = append(data, self.Rfu[:]...)

	return data
}

type EddystoneEID struct {
	EphemeralId [8]byte
}

func NewEddystoneEID(
	ephemeralId [8]byte,
) *EddystoneEID {
	return &EddystoneEID{
		EphemeralId: ephemeralId,
	}
}

func (self *EddystoneEID) Bytes() []byte {
	var frameType uint8 = 0x30
	var power uint8 = 0xf9

	var data = []byte{}
	data = append(data, frameType)
	data = append(data, power)
	data = append(data, self.EphemeralId[:]...)

	return data
}

type EddystoneURL struct {
	Prefix byte
	Url [17]byte
}

func NewEddystoneURL(
	prefix byte,
	url [17]byte,
) *EddystoneURL {
	return &EddystoneURL{
		Prefix: prefix,
		Url: url,
	}
}

func (self *EddystoneURL) Bytes() []byte {
	var frameType uint8 = 0x10
	var power uint8 = 0xf9

	var data = []byte{}
	data = append(data, frameType)
	data = append(data, power)
	data = append(data, self.Prefix)
	data = append(data, self.Url[:]...)

	return data
}

type EddystoneTLM struct {
	Version uint8
	Battery uint16
	Temperature int16
	PduCount uint32
	Time uint32
}

func NewEddystoneTLM(
	version uint8,
	battery uint16,
	temperature int16,
	pduCount uint32,
	time uint32,
) *EddystoneTLM {
	return &EddystoneTLM{
		Version: version,
		Battery: battery,
		Temperature: temperature,
		PduCount: pduCount,
		Time: time,
	}
}

func (self *EddystoneTLM) Bytes() []byte {
	var frameType uint8 = 0x20

	var data = []byte{}
	data = append(data, frameType)
	data = append(data, self.Version)
	data = binary.BigEndian.AppendUint16(data, self.Battery)
	data = binary.BigEndian.AppendUint32(data, uint32(self.Temperature))
	data = binary.BigEndian.AppendUint32(data, self.PduCount)
	data = binary.BigEndian.AppendUint32(data, self.Time)

	return data
}

type Eddystone[T EddystoneFrame] struct {
	EddystoneFrame EddystoneFrame
}

func NewEddystone[T EddystoneFrame](eddystoneStone T) *Eddystone[T] {
	return &Eddystone[T] {
		EddystoneFrame: eddystoneStone,
	}
}

func (self *Eddystone[T]) ConvertToBluetoothAdvertisement(name string) *bluetooth.BluetoothAdvertisement {
	var bluetoothAdvertisement = &bluetooth.BluetoothAdvertisement{
		AdvertiseSettings: bluetooth.NewBleAdvertisingSettings(
			name,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		),
	}

	bluetoothAdvertisement.AdvertiseData = &bluetooth.BleAdvertisingData{
		ServiceData: []*bluetooth.ServiceDataModel{
			{
				Uuid: uuid.MustParse("0000FEAA-0000-1000-8000-00805F9B34FB"),
				Data: self.EddystoneFrame.Bytes(),
			},
		},
	}

	return bluetoothAdvertisement
}


type AltBeacon struct {
	Uuid uuid.UUID
	ManufacturerId uint16
	AdditionalData [4]byte
	ManufacturerReserved byte
}

func NewAltBeacon(uuid uuid.UUID, manufacturerId uint16, additionalData [4]byte, manufacturerReserved byte) *AltBeacon {
	return &AltBeacon { uuid,
		manufacturerId,
		additionalData,
		manufacturerReserved,
	}
}

func (self *AltBeacon) ConvertToBluetoothAdvertisement(name string) *bluetooth.BluetoothAdvertisement {
	var bluetoothAdvertisement = &bluetooth.BluetoothAdvertisement{
		AdvertiseSettings: bluetooth.NewBleAdvertisingSettings(
			name,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		),
	}

	var manufacturerData = []byte {
		0xbe, 0xac,
	}

	uuid, err := self.Uuid.MarshalBinary()
	if err != nil {
		panic("Failed to convert UUID" + err.Error())
	}
	manufacturerData = append(manufacturerData, uuid...)
	manufacturerData = append(manufacturerData, self.AdditionalData[:]...)
	manufacturerData = append(manufacturerData, 0xf9)
	manufacturerData = append(manufacturerData, self.ManufacturerReserved)

	bluetoothAdvertisement.AdvertiseData = &bluetooth.BleAdvertisingData{
		ManufacturerData: []*bluetooth.ManufacturerDataModel{
			{
				Id: self.ManufacturerId,
				Data: manufacturerData,
			},
		},
	}

	return bluetoothAdvertisement
}