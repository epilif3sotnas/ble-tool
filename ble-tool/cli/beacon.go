package cli


import(
	// std
	"encoding/hex"

	// external
	"github.com/google/uuid"
)


type IBeacon struct {
	Uuid string  		`name:"uuid" help:"iBeacon UUID field."`
	Major int16			`name:"major" help:"iBeacon Major field (int16 value)."`
	Minor int16			`name:"minor" help:"iBeacon Minor field (int16 value)."`
}

func (b *IBeacon) Run() error {
	_, err := uuid.Parse(b.Uuid)
	return err
}

type Eddystone struct {
	Uid EddystoneUID  			`cmd:"uid" help:"Start advertising process with EddystoneUID."`
	Url EddystoneURL			`cmd:"url" help:"Start advertising process with EddystoneURL."`
	Tlm EddystoneTLM			`cmd:"tlm" help:"Start advertising process with EddystoneTLM."`
	Eid EddystoneEID			`cmd:"eid" help:"Start advertising process with EddystoneEID."`
}

type EddystoneUID struct {
	NamespaceId string  			`name:"namespace-id" help:"EddystoneUID namespace id field (hex string with 10 bytes)."`
	InstanceId string				`name:"instance-id" help:"EddystoneUID instance id field (hex string with 6 bytes)."`
	Rfu string						`name:"rfu" help:"EddystoneUID RFU field (hex string with 2 bytes)."`
}

func (b *EddystoneUID) Run() error {
	namespaceId, err := hex.DecodeString(b.NamespaceId)
	if err != nil {
		return err
	}
	if len(namespaceId) > 10 {
		return NewHexStringTooLong("namespace-id", len(namespaceId), 10)
	}
	if len(namespaceId) < 10 {
		return NewHexStringTooShort("namespace-id", len(namespaceId), 10)
	}

	instanceId, err := hex.DecodeString(b.InstanceId)
	if err != nil {
		return err
	}
	if len(instanceId) > 6 {
		return NewHexStringTooLong("instace-id", len(instanceId), 6)
	}
	if len(instanceId) < 6 {
		return NewHexStringTooShort("instance-id", len(instanceId), 6)
	}

	rfu, err := hex.DecodeString(b.Rfu)
	if err != nil {
		return err
	}
	if len(rfu) > 2 {
		return NewHexStringTooLong("rfu", len(rfu), 2)
	}
	if len(rfu) < 2 {
		return NewHexStringTooShort("rfu", len(rfu), 2)
	}

	return nil
}

type EddystoneURL struct {
	Prefix string  						`name:"prefix" help:"EddystoneURL prefix field (hex string with 1 bytes)."`
	Url string							`name:"url" help:"EddystoneURL URL field (hex string with 17 bytes)."`
}

func (b *EddystoneURL) Run() error {
	prefix, err := hex.DecodeString(b.Prefix)
	if err != nil {
		return err
	}
	if len(prefix) > 1 {
		return NewHexStringTooLong("prefix", len(prefix), 1)
	}
	if len(prefix) < 1 {
		return NewHexStringTooShort("prefix", len(prefix), 1)
	}

	url, err := hex.DecodeString(b.Url)
	if err != nil {
		return err
	}
	if len(url) > 17 {
		return NewHexStringTooLong("url", len(url), 17)
	}
	if len(url) < 17 {
		return NewHexStringTooShort("url", len(url), 17)
	}

	return nil
}

type EddystoneTLM struct {
	Version uint8  						`name:"version" help:"EddystoneTLM version field."`
	Battery uint16						`name:"battery" help:"EddystoneTLM battery field."`
	Temperature int16					`name:"temperature" help:"EddystoneTLM temperature field."`
	PduCount uint32						`name:"pdu-count" help:"EddystoneTLM PDU count field."`
	Time uint32							`name:"time" help:"EddystoneTLM time field."`
}

func (b *EddystoneTLM) Run() error {
	return nil
}

type EddystoneEID struct {
	Ephemeral string					`name:"ephemeral" help:"EddystoneEID ephemeral field (hex string with 8 bytes)."`
}

func (b *EddystoneEID) Run() error {
	ephemeral, err := hex.DecodeString(b.Ephemeral)
	if err != nil {
		return err
	}
	if len(ephemeral) > 8 {
		return NewHexStringTooLong("ephemeral", len(ephemeral), 8)
	}
	if len(ephemeral) < 8 {
		return NewHexStringTooShort("ephemeral", len(ephemeral), 8)
	}

	return nil
}

type AltBeacon struct {
	Uuid string  						`name:"uuid" help:"AltBeacon UUID field."`
	ManufacturerId uint16				`name:"manufacturer-id" help:"AltBeacon Manufacturer Id field (uint16 value)."`
	AdditionalData string				`name:"additional-data" help:"AltBeacon Additional Data field (hex string with 4 bytes)."`
	ManufacturerReserved string			`name:"manufacturer-reserved" help:"AltBeacon Manufacturer Reserved field (hex string with 1 byte)."`
}

func (b *AltBeacon) Run() error {
	_, err := uuid.Parse(b.Uuid)
	if err != nil {
		return err
	}

	additionalData, err := hex.DecodeString(b.AdditionalData)
	if err != nil {
		return err
	}
	if len(additionalData) > 4 {
		return NewHexStringTooLong("additional-data", len(additionalData), 4)
	}
	if len(additionalData) < 4 {
		return NewHexStringTooShort("additional-data", len(additionalData), 4)
	}

	manufacturerReserved, err := hex.DecodeString(b.ManufacturerReserved)
	if err != nil {
		return err
	}
	if len(manufacturerReserved) > 1 {
		return NewHexStringTooLong("manufacturer-reserved", len(manufacturerReserved), 1)
	}
	if len(manufacturerReserved) < 1 {
		return NewHexStringTooShort("manufacturer-reserved", len(manufacturerReserved), 1)
	}

	return nil
}