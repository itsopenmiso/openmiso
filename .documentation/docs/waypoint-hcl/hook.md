---
layout: docs
page_title: hook - waypoint.hcl
description: |-
  The `hook` stanza configures hooks that are executed before or after operations. This can be useful to do things such as perform a security scan on an image, run database migrations on a deploy, etc.
---

# `hook` Stanza

<Placement
  groups={[
    ['app', 'build', 'hook'],
    ['app', 'build', 'registry', 'hook'],
    ['app', 'deploy', 'hook'],
    ['app', 'release', 'hook'],
  ]}
/>

The `hook` stanza configures [hooks](../docs/lifecycle/hooks) that are
executed before or after operations. This can be useful to do things such as
perform a security scan on an image, run database migrations on a deploy, etc.

Multiple `hook` stanzas can be specified. This will execute multiple
hooks for an operation. The behavior with multiple hooks is
[documented on the hooks lifecycle page](../docs/lifecycle/hooks#execution-order).

```hcl
app "frontend" {
  build {
    # ...

    hook {
      when    = "before"
      command = ["./validate-creds.sh"]
    }
  }

  # ...
}
```

## `hook` Parameters

### Required

- `when` `(string)` - When the hook should be executed. Either "before" or "after".

- `command` `(array<string>)` - The command to execute. The first element of
  the list is the command to execute and each remainder is an argument. By
  specifying this as a list of strings, you do not need to worry about
  escaping arguments.

### Optional

- `on_failure` `(string: "fail")` - Behavior when the hook fails. If this is
  "continue" then failures are ignored. Otherwise, a failure cases the entire
  operation to fail. See [failure behavior](../docs/lifecycle/hooks#failure-behavior).
