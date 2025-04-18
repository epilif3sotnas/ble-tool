package bluetooth

import (
	// std
	"strings"

	// external
	"tinygo.org/x/bluetooth"
)


var devices map[string]bool = make(map[string]bool)
var adapter = bluetooth.DefaultAdapter


type BeaconType int
const (
	EmptyBeacon BeaconType = iota
	AltBeacon
	Eddystone
	IBeacon
)
var beaconType = []string{ "", "AltBeacon", "Eddystone", "iBeacon" }

type BeaconData interface {
	ConvertToBluetoothAdvertisement(name string) *BluetoothAdvertisement
}


type BleHandler struct {}


func NewBleHandler() *BleHandler {
	return &BleHandler{}
}

func (self *BleHandler) Enable() {
	self.must("enable BLE stack", adapter.Enable())
}

func (self *BleHandler) Scan(scanSettings *ScanSettings) {
	// Start scanning.
	println("Scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		if (device.LocalName() == "") {
			return
		}

		if scanSettings.Filter != "" && !strings.Contains(device.LocalName(), scanSettings.Filter) {
			return
		}

		if (self.contains(devices, device.LocalName())) {
			return
		}

		devices[device.LocalName()] = true
		var keys []string
		for key := range devices {
			keys = append(keys, key)
		}

		println("[", strings.Join(keys, ", "), "]")
		println("\n-----------------", device.LocalName(), "-------------------")
		println("Name:", device.LocalName())
		println("RSSI:", device.RSSI)
		println("MAC Address:", device.Address.String())

		beacon := self.beaconType(device.AdvertisementPayload.ManufacturerData(), device.AdvertisementPayload.ServiceData())

		isBeacon := false
		if beacon == EmptyBeacon {
			isBeacon = false
		} else {
			isBeacon = true
		}
		println("Is Beacon?:", isBeacon)
		if isBeacon {
			println("Beacon:", beaconType[beacon])
		}
		println("------------------------------------")
	})
	self.must("start scan", err)
}

func (self *BleHandler) beaconType(manufacturerData []bluetooth.ManufacturerDataElement, serviceData []bluetooth.ServiceDataElement) BeaconType {
	if len(manufacturerData) == 0 && len(serviceData) == 0 {
		return EmptyBeacon
	}

	for _, element := range manufacturerData {
		if (len(element.Data) <= 2) {
			return EmptyBeacon
		}

		if element.CompanyID == 0x004c && element.Data[0] == 0x02 {
			return IBeacon
		} else if element.Data[0] == 0xbe && element.Data[1] == 0xac {
			return AltBeacon
		}
	}

	for _, element := range serviceData {
		if element.UUID.Get16Bit() == 0xfeaa {
			return Eddystone
		}
	}

	return EmptyBeacon
}

func (self *BleHandler) Advertise(advertisementSettings *AdvertisementSettings) {
	bleAdv, err := ConvertToBluetoothAdvertisement(advertisementSettings.Data.ConvertToBluetoothAdvertisement(advertisementSettings.Name))
	if err != nil {
		panic("Failed to convert BluetoothAdvertisement to TinyGo Advertisement - " + err.Error())
	}

	self.must("start adv", bleAdv.Start())
	println("advertising...")
}

func (self *BleHandler) must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}

func (self *BleHandler) contains(set map[string]bool, item string) bool {
    _, ok := set[item]
    return ok
}