

resource "fauna_collection" "users" {
  name = "users"
}

resource "fauna_index" "user_by_email" {
  sources = [fauna_collection.users.name]
  name    = "user_by_email"
}

