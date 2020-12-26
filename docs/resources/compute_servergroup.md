---
subcategory: "Elastic Cloud Server (ECS)"
---

# sbercloud\_compute\_servergroup

Manages Server Group resource within Sbercloud.

## Example Usage

```hcl
resource "sbercloud_compute_servergroup" "test-sg" {
  name     = "my-sg"
  policies = ["anti-affinity"]
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the server group resource. If omitted, the provider-level region will be used. Changing this creates a new server group resource.

* `name` - (Required, String, ForceNew) A unique name for the server group. Changing this creates
    a new server group.

* `policies` - (Required, List, ForceNew) The set of policies for the server group. Only two
    policies are available right now, and both are mutually exclusive. Possible values are "affinity" and "anti-affinity". 
    "affinity": All instances/servers launched in this group will be hosted on the same compute node.
    "anti-affinity": All instances/servers launched in this group will be hosted on different compute nodes.
    Changing this creates a new server group.

* `value_specs` - (Optional, Map, ForceNew) Map of additional options.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

* `members` - The instances that are part of this server group.

## Import

Server Groups can be imported using the `id`, e.g.

```
$ terraform import sbercloud_compute_servergroup.test-sg 1bc30ee9-9d5b-4c30-bdd5-7f1e663f5edf
```
