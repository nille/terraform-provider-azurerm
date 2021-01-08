## 2.42.0 (January 08, 2021)

BREAKING CHANGES

* `azurerm_key_vault` - the field `soft_delete_enabled` is now defaulted to `true` to match the breaking change in the Azure API where Key Vaults now have Soft Delete enabled by default, which cannot be disabled. This property is now non-functional, defaults to `true` and will be removed in version 3.0 of the Azure Provider. ([#10088](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10088))
* `azurerm_key_vault` - the field `soft_delete_retention_days` is now defaulted to `90` days to match the Azure API behaviour, as the Azure API does not return a value for this field when not explicitly configured, so defaulting this removes a diff with `0`. ([#10088](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10088))

FEATURES:

* **New Data Source:** `azurerm_eventgrid_domain_topic` ([#10050](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10050))
* **New Data Source:** `azurerm_ssh_public_key` ([#9842](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9842))
* **New Resource:** `azurerm_data_factory_linked_service_synapse` ([#9928](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9928))
* **New Resource:** `azurerm_disk_access` ([#9889](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9889))
* **New Resource:** `azurerm_media_streaming_locator` ([#9992](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9992))
* **New Resource:** `azurerm_sentinel_alert_rule_fusion` ([#9829](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9829))
* **New Resource:** `azurerm_ssh_public_key` ([#9842](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9842))

IMPROVEMENTS:

* batch: updating to API version `2020-03-01` ([#10036](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10036))
* dependencies: upgrading to `v49.2.0` of `github.com/Azure/azure-sdk-for-go` ([#10042](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10042))
* dependencies: upgrading to `v0.15.1` of `github.com/tombuildsstuff/giovanni` ([#10035](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10035))
* Data Source: `azurerm_hdinsight_cluster` - support for the `kafka_rest_proxy_endpoint` property ([#8064](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8064))
* Data Source: `azurerm_databricks_workspace` - support for the `tags` property ([#9933](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9933))
* Data Source: `azurerm_subscription` - support for the `tags` property ([#8064](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8064))
* `azurerm_app_service` - now supports  `detailed_error_mesage_enabled` and `failed_request_tracing_enabled ` logs settings ([#9162](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9162))
* `azurerm_app_service` - now supports  `service_tag` in `ip_restriction` blocks ([#9609](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9609))
* `azurerm_app_service_slot` - now supports  `detailed_error_mesage_enabled` and `failed_request_tracing_enabled ` logs settings ([#9162](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9162))
* `azurerm_batch_pool` support for the `public_address_provisioning_type` property ([#10036](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10036))
* `azurerm_api_management` - support `Consumption_0` for the `sku_name` property ([#6868](https://github.com/terraform-providers/terraform-provider-azurerm/issues/6868))
* `azurerm_cdn_endpoint` - only send `content_types_to_compress` and `geo_filter` to the API when actually set ([#9902](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9902))
* `azurerm_cosmosdb_mongo_collection` - correctly read back the `_id` index when mongo 3.6 ([#8690](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8690))
* `azurerm_container_group` - support for the `volume.empty_dir` property ([#9836](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9836))
* `azurerm_data_factory_linked_service_azure_file_storage` - support for the `file_share` property ([#9934](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9934))
* `azurerm_dedicated_host` - support for addtional `sku_name` values ([#9951](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9951))
* `azurerm_devspace_controller` - deprecating since new DevSpace Controllers can no longer be provisioned, this will be removed in version 3.0 of the Azure Provider ([#10049](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10049))
* `azurerm_function_app` - make `pre_warmed_instance_count` computed to use azure's default ([#9069](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9069))
* `azurerm_function_app` - now supports  `service_tag` in `ip_restriction` blocks ([#9609](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9609))
* `azurerm_hdinsight_hadoop_cluster` - allow the value `Standard_D4a_V4` for the `vm_type` property ([#10000](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10000))
* `azurerm_hdinsight_kafka_cluster` - support for the `rest_proxy` and `kafka_management_node` blocks ([#8064](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8064))
* `azurerm_key_vault` - the field `soft_delete_enabled` is now defaulted to `true` to match the Azure API behaviour where Soft Delete is force-enabled and can no longer be disabled. This field is deprecated, can be safely removed from your Terraform Configuration, and will be removed in version 3.0 of the Azure Provider. ([#10088](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10088))
* `azurerm_kubernetes_cluster` - add support for network_mode ([#8828](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8828))
* `azurerm_log_analytics_linked_service` - add validation for resource ID type ([#9932](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9932))
* `azurerm_log_analytics_linked_service` - update validation to use generated validate functions ([#9950](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9950))
* `azurerm_monitor_diagnostic_setting` - validation that `eventhub_authorization_rule_id` is a EventHub Namespace Authorization Rule ID ([#9914](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9914))
* `azurerm_monitor_diagnostic_setting` - validation that `log_analytics_workspace_id` is a Log Analytics Workspace ID ([#9914](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9914))
* `azurerm_monitor_diagnostic_setting` - validation that `storage_account_id` is a Storage Account ID ([#9914](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9914))
* `azurerm_network_security_rule` - increase allowed the number of `application_security_group` blocks allowed ([#9884](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9884))
* `azurerm_sentinel_alert_rule_ms_security_incident` - support the `alert_rule_template_guid` and `display_name_exclude_filter` properties ([#9797](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9797))
* `azurerm_sentinel_alert_rule_scheduled` - support for the `alert_rule_template_guid` property ([#9712](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9712))
* `azurerm_sentinel_alert_rule_scheduled` - support for creating incidents ([#8564](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8564))
* `azurerm_spring_cloud_app` - support the properties `https_only`, `is_public`, and `persistent_disk` ([#9957](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9957))
* `azurerm_subscription` - support for the `tags` property ([#9047](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9047))
* `azurerm_synapse_workspace` - support for the `managed_resource_group_name` property ([#10017](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10017))
* `azurerm_traffic_manager_profile` - support for the `traffic_view_enabled` property ([#10005](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10005))

BUG FIXES:

provider: will not correctly register the `Microsoft.Blueprint` and `Microsoft.HealthcareApis` RPs ([#10062](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10062))
* `azurerm_application_gateway` - allow `750` for `file_upload_limit_mb` when the sku is `WAF_v2` ([#8753](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8753))
* `azurerm_firewall_policy_rule_collection_group` - correctly validate the `network_rule_collection.destination_ports` property ([#9490](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9490))
* `azurerm_cdn_endpoint` - changing many `delivery_rule` condition `match_values` to optional ([#8850](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8850))
* `azurerm_cosmosdb_account` - always include `key_vault_id` in update requests for azure policy enginer compatibility ([#9966](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9966))
* `azurerm_cosmosdb_table` - do not call the throughput api when serverless ([#9749](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9749))
* `azurerm_key_vault` - the field `soft_delete_retention_days` is now defaulted to `90` days to match the Azure API behaviour. ([#10088](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10088))
* `azurerm_kubernetes_cluster` - parse oms `log_analytics_workspace_id` to ensure correct casing ([#9976](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9976))
* `azurerm_role_assignment` fix crash in retry logic ([#10051](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10051))
* `azurerm_storage_account` - allow hns when `account_tier` is `Premium` ([#9548](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9548))
* `azurerm_storage_share_file` - allowing files smaller than 4KB to be uploaded ([#10035](https://github.com/terraform-providers/terraform-provider-azurerm/issues/10035))

## 2.41.0 (December 17, 2020)

UPGRADE NOTES:

* `azurerm_key_vault` - Azure will be introducing a breaking change on December 31st, 2020 by force-enabling Soft Delete on all new and existing Key Vaults. To workaround this, this release of the Azure Provider still allows you to configure Soft Delete on before this date (but once this is enabled this cannot be disabled). Since new Key Vaults will automatically be provisioned using Soft Delete in the future, and existing Key Vaults will be upgraded - a future release will deprecate the `soft_delete_enabled` field and default this to true early in 2021. ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_key_vault_certificate` - Terraform will now attempt to `purge` Certificates during deletion due to the upcoming breaking change in the Azure API where Key Vaults will have soft-delete force-enabled. This can be disabled by setting the `purge_soft_delete_on_destroy` field within the `features -> keyvault` block to `false`. ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_key_vault_key` - Terraform will now attempt to `purge` Keys during deletion due to the upcoming breaking change in the Azure API where Key Vaults will have soft-delete force-enabled. This can be disabled by setting the `purge_soft_delete_on_destroy` field within the `features -> keyvault` block to `false`. ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_key_vault_secret` - Terraform will now attempt to `purge` Secrets during deletion due to the upcoming breaking change in the Azure API where Key Vaults will have soft-delete force-enabled. This can be disabled by setting the `purge_soft_delete_on_destroy` field within the `features -> keyvault` block to `false`. ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))

FEATURES:

* **New Resource:** `azurerm_eventgrid_system_topic_event_subscription` ([#9852](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9852))
* **New Resource:** `azurerm_media_job` ([#9859](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9859))
* **New Resource:** `azurerm_media_streaming_endpoint` ([#9537](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9537))
* **New Resource:** `azurerm_subnet_service_endpoint_storage_policy` ([#8966](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8966))
* **New Resource:** `azurerm_synapse_managed_private_endpoint` ([#9260](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9260))

IMPROVEMENTS:

* `azurerm_app_service` - Add support for `outbound_ip_address_list` and `possible_outbound_ip_address_list` ([#9871](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9871))
* `azurerm_disk_encryption_set` - support for updating `key_vault_key_id` ([#7913](https://github.com/terraform-providers/terraform-provider-azurerm/issues/7913))
* `azurerm_iot_time_series_insights_gen2_environment` - exposing `data_access_fqdn` ([#9848](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9848))
* `azurerm_key_vault_certificate` - performing a "purge" of the Certificate during deletion if the feature is opted-in within the `features` block, see the "Upgrade Notes" for more information ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_key_vault_key` - performing a "purge" of the Key during deletion if the feature is opted-in within the `features` block, see the "Upgrade Notes" for more information ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_key_vault_secret` - performing a "purge" of the Secret during deletion if the feature is opted-in within the `features` block, see the "Upgrade Notes" for more information ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_log_analytics_linked_service` - Add new fields `workspace_id`, `read_access_id`, and `write_access_id` ([#9410](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9410))
* `azurerm_linux_virtual_machine` - Normalise SSH keys to cover VM import cases ([#9897](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9897))
* `azurerm_subnet` - support for the `service_endpoint_policy` block ([#8966](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8966))
* `azurerm_traffic_manager_profile` - support for new field `max_return` and support for `traffic_routing_method` to be `MultiValue` ([#9487](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9487))

BUG FIXES:

* `azurerm_key_vault_certificate` - reading `dns_names` and `emails` within the `subject_alternative_names` block from the Certificate if not returned from the API ([#8631](https://github.com/terraform-providers/terraform-provider-azurerm/issues/8631))
* `azurerm_key_vault_certificate` - polling until the Certificate is fully deleted during deletion ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_key_vault_key` - polling until the Key is fully deleted during deletion ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_key_vault_secret` -  polling until the Secret is fully deleted during deletion ([#9911](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9911))
* `azurerm_log_analytics_workspace` - adding a state migration to correctly update the Resource ID ([#9853](https://github.com/terraform-providers/terraform-provider-azurerm/issues/9853))

---

For information on changes between the v2.40.0 and v2.0.0 releases, please see [the previous v2.x changelog entries](https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/CHANGELOG-v2.md).

For information on changes in version v1.44.0 and prior releases, please see [the v1.x changelog](https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/CHANGELOG-v1.md).
