---
layout: docs
page_title: 'Creating a Waypoint Plugin'
description: |-
  How Waypoint plugins work
---

# Creating Waypoint Plugins

In this guide, you will learn how to create a simple plugin that can build Go applications; we will walkthrough all the
steps needed to make a Waypoint plugin.

## Requirements

To follow this guide, you will need the following tools:

- [Go version 1.14+](https://golang.org/doc/install)
- [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/)
- [Protocol Buffer compile plugin for Go](https://grpc.io/docs/languages/go/quickstart/)
- [Example Code](https://github.com/hashicorp/waypoint-plugin-examples)

The plugin you are going create will implement the `Builder` component and will be able to compile Go applications from
source and create a compiled binary. In addition to implementing `Builder`, you will see how to implement the optional
Configurable and ConfigurableNotify interfaces, and how to define output values used in other phases of the life cycle.

## Setting Up the Project

<iframe
  width="560"
  height="315"
  src="https://www.youtube.com/embed/4zwZCpqo0j8?start=0&end=145"
  frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen
></iframe>

To scaffold the new plugin, you can use the `template` in the example code repository. Open a terminal at this location;
you will see the following structure.

```shell
├── Makefile
├── README.md
├── bin
│   └── waypoint-plugin-template
├── builder
│   ├── builder.go
│   ├── output.pb.go
│   └── output.proto
├── clone.sh
├── go.mod
├── go.sum
├── main.go
├── platform
│   ├── auth.go
│   ├── deploy.go
│   ├── destroy.go
│   ├── output.pb.go
│   └── output.proto
├── registry
│   ├── auth.go
│   ├── output.pb.go
│   ├── output.proto
│   └── registry.go
└── release
    ├── destroy.go
    ├── output.pb.go
    ├── output.proto
    └── release.go
```

The template implements all components and interfaces for Waypoint plugins; it creates a vanilla base from which you can
build your plugins. Let's copy this template and create a new plugin. To do that, you can use the `clone.sh` script in
the template folder.

`clone.sh` requires you to provide three parameters: the name of the new plugin, the destination and the Go module name; let's
create a new plugin called `gobuilder` in the current repo.

```shell
./clone.sh gobuilder ../gobuilder github.com/hashicorp/waypoint-plugin-examples/gobuilder

Created new plugin in ../gobuilder
You can build this plugin by running the following command

cd ../gobuilder && make
```

The clone script creates the new plugin at the requested path `../gobuilder`, let's change to this path to build the plugin.

```shell
cd ../gobuilder
```

The `gobuilder` folder is an exact copy of the `template,` but all the Go module paths have changed to the package name
you provided to the command. Before starting to modify the plugin, let's check you can build it. When you run the `make`
command, all the Protocol Buffers used to exchange values between plugin components, and the main plugin binary are
compiled.

```shell
Build Protos
protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./builder/output.proto
protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./registry/output.proto
protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./platform/output.proto
protoc -I . --go_opt=plugins=grpc --go_out=../../../../ ./release/output.proto

Compile Plugin
go build -o ./bin/waypoint-plugin-template ./main.go
```

If you received an error when trying to build your plugin, double-check that you have all the required pre-requisites
installed.

[Next - Registering Plugin Components](../docs/extending-waypoint/creating-plugins/main)
