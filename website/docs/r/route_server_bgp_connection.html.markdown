---
subcategory: "Network"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_route_server_bgp_connection"
description: |-
  Manages a BGP Connection for a Route Server.
---

# azurerm_route_server_bgp_connection

Manages a Bgp Connection for a Route Server

## Example Usage

```hcl
resource "azurerm_route_server_bgp_connection" "example" {
  name            = "example-rs-bgpconnection"
  route_server_id = azurerm_route_server.example.id
  peer_asn        = 65501
  peer_ip         = "169.254.21.5"
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name which should be used for this Route Server Bgp Connection. Changing this forces a new resource to be created.

* `route_server_id` - (Required) The ID of the Route Server within which this Bgp connection should be created. Changing this forces a new resource to be created.

* `peer_asn` - (Optional) The peer autonomous system number for the Route Server Bgp Connection. Changing this forces a new resource to be created.

* `peer_ip` - (Optional) The peer ip address for the Route Server Bgp Connection. Changing this forces a new resource to be created.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Route Server Bgp Connection.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/language/resources/syntax#operation-timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Route Server Bgp Connection.
* `read` - (Defaults to 5 minutes) Used when retrieving the Route Server Bgp Connection.
* `delete` - (Defaults to 30 minutes) Used when deleting the Route Server Bgp Connection.

## Import

Route Server Bgp Connections can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_route_server_bgp_connection.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Network/virtualHubs/routeServer1/bgpConnections/connection1
```
