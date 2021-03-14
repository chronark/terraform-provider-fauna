terraform {
  required_providers {
    fauna = {
      source  = "chronark/fauna"
      version = "0.0.8"
    }
  }
}

provider "fauna" {
  fauna_key = "fnAEESzh3JACBY9ci4SwvxZYy89fP1fFbAU3HtGI"
}


resource "fauna_collection" "terraform" {
  name = "terraform"
}