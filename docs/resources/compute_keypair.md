---
subcategory: "Elastic Cloud Server (ECS)"
---

# sbercloud\_compute\_keypair

Manages a keypair resource within Sbercloud.

## Example Usage

```hcl
resource "sbercloud_compute_keypair" "test-keypair" {
  name       = "my-keypair"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDy+49hbB9Ni2SttHcbJU+ngQXUhiGDVsflp2g5A3tPrBXq46kmm/nZv9JQqxlRzqtFi9eTI7OBvn2A34Y+KCfiIQwtgZQ9LF5ROKYsGkS2o9ewsX8Hghx1r0u5G3wvcwZWNctgEOapXMD0JEJZdNHCDSK8yr+btR4R8Ypg0uN+Zp0SyYX1iLif7saiBjz0zmRMmw5ctAskQZmCf/W5v/VH60fYPrBU8lJq5Pu+eizhou7nFFDxXofr2ySF8k/yuA9OnJdVF9Fbf85Z59CWNZBvcTMaAH2ALXFzPCFyCncTJtc/OVMRcxjUWU1dkBhOGQ/UnhHKcflmrtQn04eO8xDr root@terra-dev"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the keypair resource. If omitted, the provider-level region will be used. Changing this creates a new keypair resource.

* `name` - (Required, String, ForceNew) A unique name for the keypair. Changing this creates a new
    keypair.

* `public_key` - (Required, String, ForceNew) A pregenerated OpenSSH-formatted public key.
    Changing this creates a new keypair.

* `value_specs` - (Optional, Map, ForceNew) Map of additional options.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Specifies a resource ID in UUID format.

## Import

Keypairs can be imported using the `name`, e.g.

```
$ terraform import sbercloud_compute_keypair.my-keypair test-keypair
```
