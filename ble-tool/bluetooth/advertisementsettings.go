package bluetooth


type AdvertisementSettings struct {
	Name string
	Data BeaconData
}

func NewAdvertisementSettings(name string, data BeaconData) *AdvertisementSettings {
	return &AdvertisementSettings{
		Name: name,
		Data: data,
	}
}