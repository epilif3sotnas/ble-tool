package cli

import(
	// std
	"fmt"
)


type Global struct {
	Version VersionGlobal `cmd:"" help:"Print version information and quit"`
}

type VersionGlobal struct {}

func (v *VersionGlobal) Run() error {
	fmt.Println("ble-tool", version)
	return nil
}