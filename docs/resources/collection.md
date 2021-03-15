---
page_title: "fauna_collection Resource - terraform-provider-fauna"
subcategory: ""
description: |-
  Manipulate collections, watch out when removing ttl_days or history_days from a terraform config. Currently they will not reset to the fauna defaults. You need to set them manually.
---

# Resource `fauna_collection`

Manipulate collections, watch out when removing `ttl_days` or `history_days` from a terraform config. Currently they will not reset to the fauna defaults. You need to set them manually.

## Example Usage

```terraform
resource "fauna_collection" "minimal" {
  name = "users"
}


resource "fauna_collection" "full_configuration" {
  name         = "users"
  ttl_days     = 90
  history_days = 30

}
```

## Schema

### Required

- **name** (String)

### Optional

- **history_days** (Number) Will not reset to the fauna defaults when you delete it. Please manually set it to `30`
- **ttl_days** (Number) Will not reset to the fauna defaults when you delete it. Please manually set it to `0`

### Read-only

- **id** (String) The id of this resource.
- **ts** (Number)


