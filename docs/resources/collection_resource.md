---
page_title: "collection_resource Resource - terraform-provider-fauna"
subcategory: ""
description: |-
  Manipulate collections
---

# Resource `collection_resource`

Manipulate collections

## Example Usage

```terraform
resource "collection_resource" "users" {
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


