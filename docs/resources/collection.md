---
page_title: "fauna_collection Resource - terraform-provider-fauna"
subcategory: ""
description: |-
  Manipulate collections
---

# Resource `fauna_collection`

Manipulate collections

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

- **history_days** (Number)
- **id** (String) The ID of this resource.
- **ttl_days** (Number)

### Read-only

- **ts** (Number)


