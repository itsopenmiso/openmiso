---
layout: docs
page_title: plugin - waypoint.hcl
description: |-
  The `plugin` stanza configures a plugin that will be used by the Waypoint configuration. This can set settings such as the checksum of the plugin, the kind of features the plugin supports, etc.
---

# `plugin` Stanza

<Placement groups={[['plugin']]} />

The `plugin` stanza configures a plugin that will be used by the Waypoint
configuration. This can set settings such as the checksum of the plugin,
the kind of features the plugin supports, etc.

The `plugin` stanza is **optional.** When you use a [`use`](../docs/waypoint-hcl/use)
stanza, it implicitly defines a plugin that Waypoint will require. You do not
need to explicitly create a `plugin` configuration as well. You only need to
create a `plugin` stanza for additional plugin configuration.

Multiple `plugin` stanzas can be specified.

```hcl
plugin "my-platform" {
  checksum = "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"

  type {
    mapper   = true
    platform = true
  }
}

# ...
```

## `plugin` Parameters

### Label

The `plugin` stanza takes a label. The label of the stanza is the name
of the plugin. In the above example, the label is "my-platform"

The plugin name is significant because it defines how the plugin binary
will be named on the local system. See [plugin loading](../docs/plugins) for more
information.

### Optional

- `checksum` `(string: "")` - A SHA-256 checksum for the external plugin binary.
  This has no effect for built-in plugins.

- `type` <code>([type](../docs/waypoint-hcl/plugin#type-parameters): nil)</code> - The
  type of plugin that this is. A plugin can implement multiple types.

## `type` Parameters

The `type` stanza is used only within a `plugin` stanza. It defines what features
a plugin implements. Each field in the `type` stanza should be set to "true"
if the plugin implements that type.

### Optional

- `builder` `(bool: false)`

- `mapper` `(bool: false)`

- `platform` `(bool: false)`

- `registry` `(bool: false)`

- `releaser` `(bool: false)`

-> **Note: mappers are special.** A mapper type plugin is _always_ loaded,
because it may provide mapping functions used for other plugins to interoperate.
