package cli

import (
	// std
	"encoding/hex"

	// internal
	"ble-tool/beacons"
	"ble-tool/bluetooth"

	// external
	"github.com/google/uuid"
)


type CLI struct {
	Global
	Bluetooth
}


func NewCli() *CLI {
	return &CLI{}
}

func (c *CLI) GetScanSettings() *bluetooth.ScanSettings {
	return bluetooth.NewScanSettings(
		c.Scan.Filter,
		c.Scan.ConnectableDevices,
	)
}

func (c *CLI) GetAdvertisementsSettings(ctx KongContext) (*bluetooth.AdvertisementSettings, error) {
	cmd := ctx.Command()
	switch cmd {
		case "advertise ibeacon":
			return bluetooth.NewAdvertisementSettings(
				c.Advertise.Name,
				beacons.NewiBeacon(
					uuid.MustParse(c.Advertise.Ibeacon.Uuid),
					c.Advertise.Ibeacon.Major,
					c.Advertise.Ibeacon.Minor,
				),
			), nil

		case "advertise altbeacon":
			additionalData, _ := hex.DecodeString(c.Advertise.Altbeacon.AdditionalData)
			manufacturerReserved, _ := hex.DecodeString(c.Advertise.Altbeacon.ManufacturerReserved)

			return bluetooth.NewAdvertisementSettings(
				c.Advertise.Name,
				beacons.NewAltBeacon(
					uuid.MustParse(c.Advertise.Altbeacon.Uuid),
					c.Advertise.Altbeacon.ManufacturerId,
					[4]byte(additionalData), manufacturerReserved[0],
				),
			), nil

		case "advertise eddystone uid":
			namespaceId, _ := hex.DecodeString(c.Advertise.Eddystone.Uid.NamespaceId)
			instanceId, _ := hex.DecodeString(c.Advertise.Eddystone.Uid.InstanceId)
			rfu, _ := hex.DecodeString(c.Advertise.Eddystone.Uid.Rfu)

			return bluetooth.NewAdvertisementSettings(
				c.Advertise.Name,
				beacons.NewEddystone(
					beacons.NewEddystoneUID(
						[10]byte(namespaceId),
						[6]byte(instanceId),
						[2]byte(rfu),
					),
				),
			), nil

		case "advertise eddystone url":
			prefix, _ := hex.DecodeString(c.Advertise.Eddystone.Url.Prefix)
			url, _ := hex.DecodeString(c.Advertise.Eddystone.Url.Url)

			return bluetooth.NewAdvertisementSettings(
				c.Advertise.Name,
				beacons.NewEddystone(
					beacons.NewEddystoneURL(
						prefix[0],
						[17]byte(url),
					),
				),
			), nil

		case "advertise eddystone tlm":
			return bluetooth.NewAdvertisementSettings(
				c.Advertise.Name,
				beacons.NewEddystone(
					beacons.NewEddystoneTLM(
						c.Advertise.Eddystone.Tlm.Version,
						c.Advertise.Eddystone.Tlm.Battery,
						c.Advertise.Eddystone.Tlm.Temperature,
						c.Advertise.Eddystone.Tlm.PduCount,
						c.Advertise.Eddystone.Tlm.Time,
					),
				),
			), nil

		case "advertise eddystone eid":
			ephemeral, _ := hex.DecodeString(c.Advertise.Eddystone.Eid.Ephemeral)

			return bluetooth.NewAdvertisementSettings(
				c.Advertise.Name,
				beacons.NewEddystone(
					beacons.NewEddystoneEID(
						[8]byte(ephemeral),
					),
				),
			), nil

		default:
			return nil, NewCommandNotSupported(cmd)
	}
}