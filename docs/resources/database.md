---
page_title: "fauna_database Resource - terraform-provider-fauna"
subcategory: ""
description: |-
  Create, Read, Update or Delete databases
---

# Resource `fauna_database`

Create, Read, Update or Delete databases



## Schema

### Required

- **name** (String) The name of a database. Databases cannot be named any of the following reserved words: `events`, `set`, `self`, `documents`, or `_`.

### Read-only

- **id** (String) The id of this resource.
- **ref** (String) A Reference to the database that was created.
- **ts** (Number) The timestamp, with microsecond resolution, associated with the creation of the database.


