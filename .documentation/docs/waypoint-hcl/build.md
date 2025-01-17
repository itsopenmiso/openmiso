---
layout: docs
page_title: build - waypoint.hcl
description: |-
  The `build` stanza describes how an application is built. During the build operation, Waypoint converts application source into an artifact and optionally pushes it to a remote registry.
---

# `build` Stanza

<Placement groups={[['app', 'build']]} />

The `build` stanza describes how an application is built. Waypoint converts
application source code into an artifact. Waypoint can store the image locally,
or push it to a remote registry. Read more about the
[build phase of the application lifecycle here](../docs/lifecycle/build).

Exactly one `build` stanza is required within an `app` stanza.

```hcl
app "frontend" {
  build {
    use "docker" {}

    registry {
      # ...
    }

    hook {
      # ...
    }
  }

  # ...
}
```

## `build` Parameters

### Required

- `use` <code>([use][use])</code> - The plugin to use for building the
  application. For example, the "docker" plugin will use `docker build` to
  build your application.

### Optional

- `hook` <code>([hook][hook]: nil)</code> - [Hooks](../docs/lifecycle/hooks)
  to execute before and/or after the build.

- `registry` <code>([registry][registry]: nil)</code> - A registry to
  push the built artifact to. If this isn't specified, the artifact isn't
  pushed to any registry and it is assumed that the deployment can access
  the build result.

[hook]: /waypoint/docs/waypoint-hcl/hook 'Hook Stanza'
[registry]: /waypoint/docs/waypoint-hcl/registry 'Registry Stanza'
[use]: /waypoint/docs/waypoint-hcl/use 'Use Stanza'
