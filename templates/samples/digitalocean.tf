
provider "digitalocean" {
  # You need to set this in your .bashrc  # export DIGITALOCEAN_TOKEN="Your API TOKEN"  #
  version = "= 1.19.0"
}

terraform {
  required_version = ">= 0.12"
}
