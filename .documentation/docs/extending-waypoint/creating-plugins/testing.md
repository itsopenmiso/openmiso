---
layout: docs
page_title: 'Testing the Plugin'
description: |-
  How Waypoint plugins work
---

# Testing the Plugin

Video tutorial below:

<iframe
  width="560"
  height="315"
  src="https://www.youtube.com/embed/4zwZCpqo0j8?start=1396"
  frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen
></iframe>

Let's build our example application; first, we need to run `waypoint init` to set up the application.

```shell-session
$ waypoint init
✓ Configuration file appears valid
✓ Local mode initialized successfully
✓ Project "guides" and all apps are registered with the server.
✓ Plugins loaded and configured successfully
✓ Authentication requirements appear satisfied.

Project initialized!

You may now call 'waypoint up' to deploy your project or
commands such as 'waypoint build' to perform steps individually.
```

~> **If you see an error that the plugin could not be found**, please
ensure you have installed the plugin binary to the proper location. For
additional help steps, see the [troubleshooting plugin installation](../docs/plugins#troubleshooting)
section.

```shell-session
$ waypoint build
✓ Application built successfully
```

If you look in the application folder, you will see a new binary called `app` which is the product of your plugin.

```shell-session
$ tree .
.
├── app
├── data.db
├── main.go
└── waypoint.hcl
```

Let's run this; you will see the message "Hello Waypoint" that you wrote in the `main.go`.

```shell-session
$ ./app
Hello Waypoint
```

## Summary

You have now successfully created a simple `Build` plugin for Waypoint. Creating other plugin types for
`ReleaseManager`, or `Platform` follow similar concepts. Details on the interfaces for implementing these other plugins
can be found in the [Plugin Components](../docs/extending-waypoint/plugin-interfaces/configurable) section of the documentation. The
Extending Waypoint section also contains other reference documentation related to plugin creation.

A full example of the plugin created in this guide can be found at the following location:
[https://github.com/hashicorp/waypoint-plugin-examples/tree/main/plugins/gobuilder_final](https://github.com/hashicorp/waypoint-plugin-examples/tree/main/plugins/gobuilder_final)
