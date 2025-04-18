package cli_test

import (
	// std
	"testing"

	// internal
	"ble-tool/beacons"
	"ble-tool/bluetooth"
	"ble-tool/cli"

	// external
	"github.com/google/uuid"
	"github.com/petergtz/pegomock/v4"
	"github.com/stretchr/testify/assert"
)


// Mocks to run this test
//go:generate pegomock generate ble-tool/cli KongContext


func TestGetScanSettingsWithDataValid(t *testing.T) {
	cliTest := cli.NewCli()
	scan := cli.ScanBluetooth { Filter: "SCOOTER", ConnectableDevices: true }
	cliTest.Scan = scan

	expected := bluetooth.ScanSettings { Filter: scan.Filter, ConnectableDevices: scan.ConnectableDevices }
	actual := cliTest.GetScanSettings()

	assert.Equalf(t, expected.Filter, actual.Filter, "Verification of Filter")
	assert.Equalf(t, expected.ConnectableDevices, actual.ConnectableDevices, "Verification of Connectable Devices")
}

func TestGetScanSettingsWithDataNotValid(t *testing.T) {
	cliTest := cli.NewCli()
	scan := cli.ScanBluetooth {}
	cliTest.Scan = scan

	actual := cliTest.GetScanSettings()

	assert.Emptyf(t, actual.Filter, "Verification of Filter Empty")
	assert.Emptyf(t, actual.ConnectableDevices, "Verification of Connectable Devices Empty")
}

func TestGetAdvertisementsSettingsValidiBeacon(t *testing.T) {
	mockedKongContext := NewMockKongContext(pegomock.WithT(t))
	pegomock.When(mockedKongContext.Command()).ThenReturn("advertise ibeacon")


	cliTest := cli.NewCli()
	advertisement := cli.AdvertiseBluetooth {
		Name: "TestName",
		Ibeacon: cli.IBeacon {
			Uuid: "c55dcf88-a7d4-4fdc-9559-9f1f564f6fb3",
			Major: 10,
			Minor: 10,
		},
	}
	cliTest.Advertise = advertisement

	expected := &bluetooth.AdvertisementSettings {
		Name: "TestName",
		Data: beacons.NewiBeacon(
			uuid.MustParse(advertisement.Ibeacon.Uuid),
			10,
			10,
		),
	}
	actual, err := cliTest.GetAdvertisementsSettings(mockedKongContext)

	assert.Nil(t, err)
	assert.Equalf(t, expected, actual, "Verification of Advertsiment Settings.")
}

func TestGetAdvertisementsSettingsValidAltBeacon(t *testing.T) {
	mockedKongContext := NewMockKongContext(pegomock.WithT(t))
	pegomock.When(mockedKongContext.Command()).ThenReturn("advertise altbeacon")


	cliTest := cli.NewCli()
	advertisement := cli.AdvertiseBluetooth {
		Name: "TestName",
		Altbeacon: cli.AltBeacon {
			Uuid: "c55dcf88-a7d4-4fdc-9559-9f1f564f6fb3",
			ManufacturerId: 10,
			AdditionalData: "aabbccdd",
			ManufacturerReserved: "ad",
		},
	}
	cliTest.Advertise = advertisement

	expected := &bluetooth.AdvertisementSettings {
		Name: "TestName",
		Data: beacons.NewAltBeacon(
			uuid.MustParse(advertisement.Altbeacon.Uuid),
			10,
			[4]byte{170, 187, 204, 221},
			173,
		),
	}
	actual, err := cliTest.GetAdvertisementsSettings(mockedKongContext)

	assert.Nil(t, err)
	assert.Equalf(t, expected, actual, "Verification of Advertsiment Settings.")
}

func TestGetAdvertisementsSettingsInvalidCommand(t *testing.T) {
	mockedKongContext := NewMockKongContext(pegomock.WithT(t))
	pegomock.When(mockedKongContext.Command()).ThenReturn("invalid command")


	cliTest := cli.NewCli()
	_, err := cliTest.GetAdvertisementsSettings(mockedKongContext)

	assert.NotNil(t, err)
}