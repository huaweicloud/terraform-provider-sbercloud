---
subcategory: "Elastic Load Balance (ELB)"
---

# sbercloud\_lb\_member

Manages an ELB member resource within SberCloud.
## Example Usage

```hcl
resource "sbercloud_lb_member" "member_1" {
  address       = "192.168.199.23"
  protocol_port = 8080
  pool_id       = var.pool_id
  subnet_id     = var.subnet_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the ELB member resource.
    If omitted, the the provider-level region will be used.
    Changing this creates a new member.

* `pool_id` - (Required, String, ForceNew) The id of the pool that this member will be
    assigned to.

* `subnet_id` - (Required, String, ForceNew) The subnet in which to access the member

* `name` - (Optional, String) Human-readable name for the member.

* `tenant_id` - (Optional, String, ForceNew) Required for admins. The UUID of the tenant who owns
    the member.  Only administrative users can specify a tenant UUID
    other than their own. Changing this creates a new member.

* `address` - (Required, String, ForceNew) The IP address of the member to receive traffic from
    the load balancer. Changing this creates a new member.

* `protocol_port` - (Required, Int, ForceNew) The port on which to listen for client traffic.
    Changing this creates a new member.

* `weight` - (Optional, Int)  A positive integer value that indicates the relative
    portion of traffic that this member should receive from the pool. For
    example, a member with a weight of 10 receives five times as much traffic
    as a member with a weight of 2.

* `admin_state_up` - (Optional, Bool) The administrative state of the member.
    A valid value is true (UP) or false (DOWN).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique ID for the member.

## Timeouts
This resource provides the following timeouts configuration options:
- `create` - Default is 10 minute.
- `update` - Default is 10 minute.
- `delete` - Default is 10 minute.
