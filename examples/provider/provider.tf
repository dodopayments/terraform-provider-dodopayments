terraform {
  required_providers {
    dodopayments = {
      source = "dodopayments/dodopayments"
    }
  }
}

# The provider reads the API key from the DODO_PAYMENTS_API_KEY environment variable.
provider "dodopayments" {
  environment = "test_mode"
}
