terraform {
  required_providers {
    fauna = {
      source  = "hashicorp.com/chronark/fauna"
      version = "9000.1"
    }
  }
}

provider "fauna" {
  fauna_key = "fnAEEWZgLrACB6LzAzbAotOEPVrqCQKX1-rbedfw"
}


resource "fauna_collection" "my_collection" {
  name = "terraform15"
}