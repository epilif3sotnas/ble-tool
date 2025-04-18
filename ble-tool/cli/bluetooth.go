package cli


type Bluetooth struct {
	Advertise AdvertiseBluetooth 		`cmd:"advertise" help:"Start advertising process with some args configuration."`
	Scan ScanBluetooth					`cmd:"scan" help:"Start scan process with some args configuration."`
}

type AdvertiseBluetooth struct {
	Name string							`name:"name" short:"n" help:"Device Name of Advertisement."`
	Ibeacon IBeacon						`cmd:"ibeacon" help:"Start advertising process with iBeacon."`
	Altbeacon AltBeacon				 	`cmd:"altbeacon" help:"Start advertising process with AltBeacon."`
	Eddystone Eddystone				 	`cmd:"eddystone" help:"Start advertising process with Eddystone."`
}

type ScanBluetooth struct {
	Filter string						`name:"filter" short:"f" help:"Filter advertisements that contains this string filter."`
	ConnectableDevices bool				`name:"connectable-devices" short:"c" help:"Filter advertisements that contains only connectable devices."`
}

func (sb *ScanBluetooth) Run() error {
	return nil
}