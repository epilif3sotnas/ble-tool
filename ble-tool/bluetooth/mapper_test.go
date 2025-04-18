package bluetooth_test

import (
	// std
	"testing"
	// "time"

	// internal
	// "ble-tool/bluetooth"

	// // external
	// "github.com/stretchr/testify/assert"
	// bluetoothtiny "tinygo.org/x/bluetooth"
)



func TestConvertToBluetoothAdvertisementWithValidData(t *testing.T) {
	// advName := "TestBeacon"

	// var manufacturerId uint16 = 0x004c
	// manufacturerData := []byte {
	//     0x02, 0x15, 0xB9, 0x40, 0x7F, 0x30, 0xF5, 0xF8,
	// 	0x46, 0x6E, 0xAF, 0xF9, 0x25, 0x55, 0x6B, 0x57,
	// 	0xFE, 0x6D, 0x00, 0x01, 0x00, 0x02, 0xF9,
	// }
	// bluetoothAdvertisement := &bluetooth.BluetoothAdvertisement {
	// 	AdvertiseSettings: &bluetooth.BleAdvertisingSettings {
	// 		Name: advName,
	// 		IncludeTxPower: true,
	// 		IncludeDeviceName: true,
	// 		Interval: "250ms",
	// 		Connectable: false,
	// 		LegacyMode: true,
	// 		Scannable: true,
	// 	},
	// 	AdvertiseData: &bluetooth.BleAdvertisingData{
	// 		ManufacturerData: []*bluetooth.ManufacturerDataModel{
	// 			{
	// 				Id: manufacturerId,
	// 				Data: manufacturerData,
	// 			},
	// 		},
	// 	},
	// }

	// duration, _ := time.ParseDuration("250ms")

	// expected := bluetoothtiny.DefaultAdapter.DefaultAdvertisement()
	// expected.Configure(
	// 	bluetoothtiny.AdvertisementOptions{
	// 		LocalName: "TestBeacon",
	// 		Interval: bluetoothtiny.NewDuration(duration),
	// 		ManufacturerData: []bluetoothtiny.ManufacturerDataElement{},
	// 	},
	// )
	// actual, err := bluetooth.ConvertToBluetoothAdvertisement(bluetoothAdvertisement)

	// assert.NotNil(t, err)
	// assert.Nil(t, actual)
	// assert.Equalf(t, expected, actual, "ConvertToBluetoothAdvertisement - should get a valid Advertisement")
}