---
subcategory: "Web PubSub"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_web_pubsub"
description: |-
  Manages an Azure Web Pubsub service.
---

# azurerm_web_pubsub

Manages an Azure Web Pubsub Service.

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "terraform-webpubsub"
  location = "east us"
}

resource "azurerm_web_pubsub" "example" {
  name                = "tfex-webpubsub"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name

  sku      = "Standard_S1"
  capacity = 1

  public_network_access_enabled = false

  live_trace {
    enabled                   = true
    messaging_logs_enabled    = true
    connectivity_logs_enabled = false
  }

  identity {
    type = "SystemAssigned"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the Web Pubsub service. Changing this forces a new resource to be created.

* `resource_group_name` - (Required) The name of the resource group in which to create the Web Pubsub service. Changing
  this forces a new resource to be created.

* `location` - (Required) Specifies the supported Azure location where the Web Pubsub service exists. Changing this
  forces a new resource to be created.

* `sku` - (Required) Specifies which sku to use. Possible values are `Free_F1` and `Standard_S1`.

* `capacity` - (Optional) Specifies the number of units associated with this Web Pubsub resource. Valid values are:
  Free: `1`, Standard: `1`, `2`, `5`, `10`, `20`, `50`, `100`.

* `public_network_access_enabled` - (Optional) Whether to enable public network access? Defaults to `true`.

* `tags` - (Optional) A mapping of tags to assign to the resource.

* `live_trace` - (Optional) A `live_trace` block as defined below.

* `identity` - (Optional) An `identity` block as defined below.

* `local_auth_enabled` - (Optional) Whether to enable local auth? Defaults to `true`.

* `aad_auth_enabled` - (Optional) Whether to enable AAD auth? Defaults to `true`.

* `tls_client_cert_enabled` - (Optional)  Whether to request client certificate during TLS handshake? Defaults
  to `false`.

---

A `live_trace` block supports the following:

* `enabled` - (Optional) Whether the live trace is enabled? Defaults to `true`.

* `messaging_logs_enabled` - (Optional) Whether the log category `MessagingLogs` is enabled? Defaults to `true`

* `connectivity_logs_enabled` - (Optional) Whether the log category `ConnectivityLogs` is enabled? Defaults to `true`

* `http_request_logs_enabled` - (Optional) Whether the log category `HttpRequestLogs` is enabled? Defaults to `true`

---

An `identity` block supports the following:

* `type` - (Required) The type of identity used for the Web PubSub service. Possible values are `SystemAssigned` and `UserAssigned`. If `UserAssigned` is set, a `user_assigned_identity_id` must be set as well.

* `identity_ids` - (Optional) A list of User Assigned Identity IDs which should be assigned to this Web PubSub service.

## Attributes Reference

The following attributes are exported:

* `id` - The ID of the Web Pubsub service.

* `hostname` - The FQDN of the Web Pubsub service.

* `ip_address` - The publicly accessible IP of the Web Pubsub service.

* `public_port` - The publicly accessible port of the Web Pubsub service which is designed for browser/client use.

* `server_port` - The publicly accessible port of the Web Pubsub service which is designed for customer server side use.

* `primary_access_key` - The primary access key for the Web Pubsub service.

* `primary_connection_string` - The primary connection string for the Web Pubsub service.

* `secondary_access_key` - The secondary access key for the Web Pubsub service.

* `secondary_connection_string` - The secondary connection string for the Web Pubsub service.

## Timeouts

The `timeouts` block allows you to
specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the Web Pubsub Service.
* `update` - (Defaults to 30 minutes) Used when updating the Web Pubsub Service.
* `read` - (Defaults to 5 minutes) Used when retrieving the Web Pubsub Service.
* `delete` - (Defaults to 30 minutes) Used when deleting the Web Pubsub Service.

## Import

Web Pubsub services can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_web_pubsub.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/webPubSub/pubsub1
```

