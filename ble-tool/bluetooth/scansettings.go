package bluetooth


type ScanSettings struct {
	Filter string
	ConnectableDevices bool
}


func NewScanSettings(filter string, connectableDevices bool) *ScanSettings {
	return &ScanSettings{
		Filter: filter,
		ConnectableDevices: connectableDevices,
	}
}