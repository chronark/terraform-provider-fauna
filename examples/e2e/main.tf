terraform {
  required_providers {
    fauna = {
      source  = "hashicorp.com/chronark/fauna"
      version = "9000.1"
    }
  }
}

provider "fauna" {
  fauna_key = "fnAEESzh3JACBY9ci4SwvxZYy89fP1fFbAU3HtGI"
}


// resource "fauna_collection" "my_collection" {
//   name         = "terraform15"
// }