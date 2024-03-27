package main

import (
	"github.com/Dapacruz/nxos-cli/cmd"
	_ "github.com/Dapacruz/nxos-cli/cmd/config"
	_ "github.com/Dapacruz/nxos-cli/cmd/device"
)

func main() {
	cmd.Execute()
}
