package main

import (
	"github.com/DigitalOnUs/terraform-provider-ucs/ucs"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ucs.Provider})
}
