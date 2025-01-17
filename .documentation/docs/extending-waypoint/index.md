---
layout: docs
page_title: Extending Waypoint
description: |-
  Waypoint utilizes a server to enable collaboration with other team members, storing operation history, and enabling functionality such as logs, exec, and more.
---

# How Waypoint Plugins Work

Plugins in Waypoint are separate binaries which communicate with the Waypoint application; the plugin communicates using
gRPC. While it is theoretically possible to build a plugin in any language supported by the gRPC framework, we
recommend that the developers leverage the [Waypoint SDK](https://github.com/hashicorp/waypoint-plugin-sdk) for Go to
perform this task.

In this section, you will find reference documentation for creating plugins, if you would like to jump straight in and
start making a plugin; please check out the [Creating Waypoint Plugins Guide](../docs/extending-waypoint/creating-plugins).

## Concepts

Conceptually Waypoint Plugins can be defined using the following terms:

- `Plugin` - A Plugin is a compiled binary containing one or more Components which are executed by the Waypoint
  application
- `Component` - A Component handles part of the Waypoint application lifecycle, i.e. Build, Deploy, etc. Components are
  created by implementing various interfaces
- `Interface` - An interface implements specific behavior in a Component, i.e., Configurable, ReleaseManager
- `Output Value` - A Protocol Buffer serializable type which can be passed between components.
