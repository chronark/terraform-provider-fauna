
terraform {
  required_providers {
    fauna = {
      version = "0.2"
      source  = "hashicorp.com/chronark/fauna"
    }
  }
}

provider "fauna" {}

data "fauna_databases" "my-dbs" {
}

output "dbs" {
  value = data.fauna_databases.my-dbs.databases
}