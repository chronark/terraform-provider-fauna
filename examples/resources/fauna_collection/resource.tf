resource "fauna_collection" "minimal" {
  name = "users"
}


resource "fauna_collection" "full_configuration" {
  name         = "users"
  ttl_days     = 90
  history_days = 30

}