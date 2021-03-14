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
resource "fauna_collection" "users" {
  name = "users"
}
```

## Schema

### Required

- **name** (String)

### Optional

- **history_days** (Number) Will not reset to the fauna defaults when you delete it. Please manually set it to `30`
- **id** (String) The ID of this resource.
- **ttl_days** (Number) Will not reset to the fauna defaults when you delete it. Please manually set it to `0`

### Read-only

- **ts** (Number)


