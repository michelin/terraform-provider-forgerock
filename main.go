package main

import (
	"context"
	"flag"
	"log"

	"terraform-provider-forgerock/forgerockprovider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {

	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// TODO: Update this string with the published name of your provider.
		Address: "michelin.com/forgerock",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), forgerockprovider.New(), opts)

	if err != nil {
		log.Fatal(err.Error())
	}

}
