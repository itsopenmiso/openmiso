---
layout: docs
page_title: entrypoint - Variables - waypoint.hcl
description: |-
  The `entrypoint` variable contains important information about the Waypoint Entrypoint, such as the required environment variables that must be set for proper entrypoint behavior.
---

# `entrypoint` Variable

<Placement groups={[['app', 'deploy']]} />

The `entrypoint` variable contains important information about
the [Waypoint Entrypoint](../docs/entrypoint), such as the required
environment variables that must be set for proper entrypoint behavior.

## `entrypoint.env`

**Type: `map<string,string>`**

These are the environment variables that must be set on startup
for the entrypoint to function properly. **You do not normally need
to use this** because deployment plugins will automatically set these
environment variables in most cases.

This is useful for manual configuration in cases where the platform
may not do it automatically (for example, if you specify a custom set
of YAML files for Kubernetes).

### Example: Kubernetes

The example below shows how we can configure the Kubernetes plugin
with the required environment variables. **Note: this is not necessary
since the Kubernetes plugin automatically does this.** This is just
shown as an example for illustrative purposes.

```hcl
app "web" {
  deploy {
    use "kubernetes" {
      static_environment = merge({
        MY_CUSTOM_VAR = "hello, world"
      }, entrypoint.env)
    }
  }
}
```

This example sets some custom variables and uses the `merge` function
to merge in the Waypoint Entrypoint environment variables as well.
