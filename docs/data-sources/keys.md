---
page_title: "fauna_keys Data Source - terraform-provider-fauna"
subcategory: ""
description: |-
  A Set reference for the available authentication keys in the current database.
---

# Data Source `fauna_keys`

A Set reference for the available authentication keys in the current database.

## Example Usage

```terraform
data "fauna_keys" "all_keys" {
  database = "my_database"
}
```

## Schema

### Optional

- **database** (String) A reference to a child database. If not specified, the current database is used.

### Read-only

- **id** (String) Internal id
- **keys** (Set of Object) (see [below for nested schema](#nestedatt--keys))

<a id="nestedatt--keys"></a>
### Nested Schema for `keys`

Read-only:

- **database** (String)
- **hashed_secret** (String)
- **id** (String)
- **name** (String)
- **role** (String)
- **ts** (Number)


