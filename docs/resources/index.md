---
page_title: "fauna_index Resource - terraform-provider-fauna"
subcategory: ""
description: |-
  It is possible to rename an index by updating its name field. Renaming an index changes its Reference, but preserves inbound References to the index. Index data is not rebuilt.
  An index’s terms and values fields may not be changed. If you require such a change, the existing index must be deleted and a new one created using the new definitions for terms and/or values.
  If you update the unique field, existing duplicate items are not removed from the index.
---

# Resource `fauna_index`

It is possible to rename an index by updating its name field. Renaming an index changes its Reference, but preserves inbound References to the index. Index data is not rebuilt.
An index’s terms and values fields may not be changed. If you require such a change, the existing index must be deleted and a new one created using the new definitions for terms and/or values.
If you update the unique field, existing duplicate items are not removed from the index.



## Schema

### Required

- **name** (String) The logical name of the index. Cannot be `events`, `sets`, `self`, `documents`, or `_`.
- **source** (Set of String) An array of one or more collection names

### Optional

- **id** (String) The ID of this resource.
- **serialized** (Boolean) If `true`, writes to this index are serialized with concurrent reads and writes.
- **terms** (Block Set) An array of Term objects describing the fields that should be searchable. Indexed terms can be used to search for field values, via the `Match` function. The default is an empty Array. (see [below for nested schema](#nestedblock--terms))
- **unique** (Boolean) If `true`, maintains a unique constraint on combined `terms` and `values`.
- **values** (Block List) An array of Value objects describing the fields that should be reported in search results. The default is an empty Array. When no `values` fields are defined, search results report the indexed document’s Reference. (see [below for nested schema](#nestedblock--values))

### Read-only

- **active** (Number) When an index is added, it is immediately available for reads, but returns incomplete results until it is built. Fauna builds the index asynchronously by scanning over relevant documents. Upon completion, the index’s `active` field is set to `true`.
- **partitions** (Number) The number of sub-partitions used by this index. This value can be 8 or 1:
`1` when unique is true.
`8` when the index has no terms.
`1` in all other case.
- **ts** (Number) A timestamp when this index was created

<a id="nestedblock--terms"></a>
### Nested Schema for `terms`

Optional:

- **field** (List of String)


<a id="nestedblock--values"></a>
### Nested Schema for `values`

Required:

- **field** (List of String) The path fo the field within a document to be indexed

Optional:

- **reverse** (Boolean) Whether this field’s value should sort reversed.


