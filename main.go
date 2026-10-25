// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/dodopayments/terraform-provider-dodopayments/internal"
	"github.com/dodopayments/terraform-provider-dodopayments/internal/provider"
)

// goreleaser sets version at release time.
var version = internal.PackageVersion

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Start the provider with support for debuggers such as delve.")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.New(version), providerserver.ServeOpts{
		Address: "registry.terraform.io/dodopayments/dodopayments",
		Debug:   debug,
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
