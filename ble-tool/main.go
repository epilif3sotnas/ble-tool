package main

import (
	// std
	"strings"

	// internal
	"ble-tool/bluetooth"
	"ble-tool/cli"

	// external
	"github.com/alecthomas/kong"
)


func main() {
	cli := cli.NewCli()
	ctx := kong.Parse(
		cli,
		kong.Name("ble-tool"),
		kong.Description("BLE tool to scan and advertise."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)

	err := ctx.Run()
	ctx.FatalIfErrorf(err)

	bleHandler := bluetooth.NewBleHandler()
	bleHandler.Enable()

	cmd := strings.Split(ctx.Command(), " ")[0]

	switch cmd {
		case "scan":
			scanSettings := cli.GetScanSettings()
			go bleHandler.Scan(scanSettings)

		case "advertise":
			advertsiementSettings, _ := cli.GetAdvertisementsSettings(ctx)
			go bleHandler.Advertise(advertsiementSettings)
	}

	// Maintain main Thread active for commands that use goroutines
	switch cmd {
		case "version":
			break

		default:
			for {}
	}
}