---
layout: docs
page_title: release - waypoint.hcl
description: |-
  The `release` stanza describes how an application is released. A release activates a previously staged deployment and opens it to general traffic. This step may involve adding a deployment to a load balancer, updating DNS, configuring a service mesh, etc.
---

# `release` Stanza

<Placement groups={[['app', 'release']]} />

The `release` stanza describes how an application is released.
A release activates a previously staged deployment and opens it to general
traffic. This step may involve adding a deployment to a load balancer, updating
DNS, configuring a service mesh, etc.

Release configuration is **optional.** Please see the documentation on the
[release phase of the Waypoint lifecycle](../docs/lifecycle/release)
for the Waypoint behavior when a release phase is not configured.

```hcl
app "frontend" {
  build {
    use "docker" {}

    registry {
      # ...
    }
  }

  deploy {
    use "kubernetes" {}

    hook {
      # ...
    }
  }

  release {
    use "aws-alb" {}

    hook {
      # ...
    }
  }
# ...
}
```

## `release` Parameters

### Required

- `use` <code>([use][use])</code> - The plugin to use for release management of the
  application.

### Optional

- `hook` <code>([hook][hook]: nil)</code> - [Hooks](../docs/lifecycle/hooks)
  to execute before and/or after the release.

[hook]: /waypoint/docs/waypoint-hcl/hook 'Hook Stanza'
[use]: /waypoint/docs/waypoint-hcl/use 'Use Stanza'
