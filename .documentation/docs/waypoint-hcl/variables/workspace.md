---
layout: docs
page_title: workspace - Variables - waypoint.hcl
description: |-
  Workspace variables are used to read information about the currently active workspace.
---

# `workspace` Variables

Workspace variables are used to read information about the currently active
[workspace](../docs/workspaces).

## `workspace.name`

<Placement groups={[['* (global)']]} />

The name of the currently active workspace.

A common use case for workspaces is to model environments (staging, production, etc.).
This variable can be used to change behavior based on this workspace. In the
example below, we change the Dockerfile used based on the workspace.

Example:

```hcl
dockerfile = "${path.project}/Dockerfile.${workspace.name}"
```
