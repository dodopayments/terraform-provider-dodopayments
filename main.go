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

// scripts/format runs this directive. It writes the registry docs in docs/.
// Actions and list resources need Terraform 1.14 or later in the schema export, so tfplugindocs downloads that version.
//go:generate go tool tfplugindocs generate --provider-name dodopayments --tf-version 1.14.0

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
