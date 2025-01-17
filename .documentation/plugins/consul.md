---
layout: plugins
page_title: 'Plugin: Consul'
description: 'Source configuration from Consul KV'
---

# Consul

## consul (configsourcer)

Read configuration values from the Consul KV store.

### Required Parameters

These parameters are used in `dynamic` for sourcing [configuration values](../docs/app-config/dynamic) or [input variable values](../docs/waypoint-hcl/variables/dynamic).

#### key

The KV path to retrieve.

- Type: **string**

### Optional Parameters

These parameters are used in `dynamic` for sourcing [configuration values](../docs/app-config/dynamic) or [input variable values](../docs/waypoint-hcl/variables/dynamic).

#### allow_stale

Whether to perform a stale query for retrieving the KV data.

If not set this will default to true. It must explicitly be set to false in order to use consistent queries.

- Type: **bool**
- **Optional**

#### datacenter

The datacenter to load the KV value from.

If not specified then it will default to the plugin's global datacenter configuration. If that is also not specified then Consul will default the datacenter like it would any other request.

- Type: **string**
- **Optional**

#### namespace

The namespace to load the KV value from.

If not specified then it will default to the plugin's global namespace configuration. If that is also not specified then Consul will default the namespace like it would any other request.

- Type: **string**
- **Optional**

#### partition

The partition to load the KV value from.

If not specified then it will default to the plugin's global partition configuration. If that is also not specified then Consul will default the partition like it would any other request.

- Type: **string**
- **Optional**
