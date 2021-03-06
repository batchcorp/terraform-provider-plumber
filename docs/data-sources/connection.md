---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "plumber_connection Data Source - terraform-provider-plumber"
subcategory: ""
description: |-
  
---

# Data Source: plumber_connection

Use this data source to query a configured message bus connection in order
to use its ID within a resource configuration.

## Example Usage

---

```hcl
data "plumber_connection" "my_connection" {
  filter {
    name = "name"
    values = ["My Kafka Connection"]
  }
}
```


## Argument Reference

---

- `filter` - (Block Set) - see [below for nested schema](#nestedblock--filter)

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

- `name` - (String) - Field name to filter on
  - Allowed values:
    - `name` - The connection's name
- `values` - (List of String) - Value(s) to filter by. Wildcards `*` are supported.


## Attribute Reference

---

- `id` - (String) - Connection ID
- `name` - (String) - Connection name