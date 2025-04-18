package beacons_test


import (
	// std
	"testing"

	// internal
	"ble-tool/beacons"
	"ble-tool/bluetooth"

	// external
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


func TestConvertToBluetoothAdvertisementiBeaconValidData(t *testing.T) {
	serviceUuid := uuid.MustParse("B9407F30-F5F8-466E-AFF9-25556B57FE6D")
	ibeacon := beacons.NewiBeacon(
		serviceUuid,
		0x01,
		0x02,
	)
	advName := "TestBeacon"

	var manufacturerId uint16 = 0x004c
	manufacturerData := []byte {
	    0x02, 0x15, 0xB9, 0x40, 0x7F, 0x30, 0xF5, 0xF8,
		0x46, 0x6E, 0xAF, 0xF9, 0x25, 0x55, 0x6B, 0x57,
		0xFE, 0x6D, 0x00, 0x01, 0x00, 0x02, 0xF9,
	}
	expected := &bluetooth.BluetoothAdvertisement {
		AdvertiseSettings: &bluetooth.BleAdvertisingSettings {
			Name: advName,
			IncludeTxPower: true,
			IncludeDeviceName: true,
			Interval: "250ms",
			Connectable: false,
			LegacyMode: true,
			Scannable: true,
		},
		AdvertiseData: &bluetooth.BleAdvertisingData{
			ManufacturerData: []*bluetooth.ManufacturerDataModel{
				{
					Id: manufacturerId,
					Data: manufacturerData,
				},
			},
		},
	}
	actual := ibeacon.ConvertToBluetoothAdvertisement(advName)


	assert.Equalf(t, expected, actual, "ConvertToBluetoothAdvertisement- should be possible to convert the ibeacon data to BluetoothAdvertisment")
}

func TestConvertToBluetoothAdvertisementAltBeaconValidData(t *testing.T) {
	serviceUuid := uuid.MustParse("B9407F30-F5F8-466E-AFF9-25556B57FE6D")
	altbeacon := beacons.NewAltBeacon(
		serviceUuid,
		0x0101,
		[4]byte{0x01, 0x02, 0x03, 0x04},
		0x00,
	)
	advName := "TestBeacon"

	var manufacturerId uint16 = 0x0101
	manufacturerData := []byte {
	    0xbe, 0xac, 0xB9, 0x40, 0x7F, 0x30, 0xF5, 0xF8,
		0x46, 0x6E, 0xAF, 0xF9, 0x25, 0x55, 0x6B, 0x57,
		0xFE, 0x6D, 0x01, 0x02, 0x03, 0x04, 0xF9, 0x00,
	}
	expected := &bluetooth.BluetoothAdvertisement {
		AdvertiseSettings: &bluetooth.BleAdvertisingSettings {
			Name: advName,
			IncludeTxPower: true,
			IncludeDeviceName: true,
			Interval: "250ms",
			Connectable: false,
			LegacyMode: true,
			Scannable: true,
		},
		AdvertiseData: &bluetooth.BleAdvertisingData{
			ManufacturerData: []*bluetooth.ManufacturerDataModel{
				{
					Id: manufacturerId,
					Data: manufacturerData,
				},
			},
		},
	}
	actual := altbeacon.ConvertToBluetoothAdvertisement(advName)


	assert.Equalf(t, expected, actual, "ConvertToBluetoothAdvertisement- should be possible to convert the AltBeacon data to BluetoothAdvertisment")
}