package main

import (
	"context"
	"flag"
	"log"

	"github.com/edgedb/terraform-provider-stripe/stripe"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode,
		"debug",
		false,
		"set to true to run the provider with the debug support")
	flag.Parse()

	opts := &plugin.ServeOpts{
		ProviderFunc: stripe.Provider,
	}

	if debugMode {
		err := plugin.Debug(context.Background(), "local/edgedb/stripe", opts)
		if err != nil {
			log.Fatal(err)
		}
	}

	plugin.Serve(opts)
}
