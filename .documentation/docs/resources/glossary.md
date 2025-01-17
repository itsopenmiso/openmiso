---
layout: docs
page_title: Glossary
description: |-
  This page lists and defines technical terms used throughout Waypoint. Words such as "project", "app", etc. can be overloaded in the technical community, so this page attempts to clarify their meaning in the context of Waypoint.
---

# Waypoint Glossary

This page lists and defines technical terms used throughout Waypoint. Words
such as "project", "app", etc. can be overloaded in the technical community,
so this page attempts to clarify their meaning in the context of Waypoint.

The easiest way to navigate this page is to use the "Jump to Section"
dropdown under the title above or use your browser's search functionality
(Cmd-F or Ctrl-F).

## Project

A project in Waypoint is represented by a single "waypoint.hcl" file
and is a 1:1 mapping to a VCS repository (if one exists). A project can
be comprised of one or more applications.

The purpose of the "project" is to give Waypoint a single name to refer
to a repository or collection of applications. Some features are used at
the project scope, such as setting [application configuration](../docs/app-config).

## Application (or "App")

An application is defined using the `app` stanza in a "waypoint.hcl" and
represents a unit that is deployed. Multiple applications may exist in
a single project, which may represent a project that is comprised of
separate deployable units. For example: a backend, a frontend, a service
worker, etc. may each be a separate application.

## Parameter

A "parameter" is the term used to describe any configuration object.

## Stanza

A "stanza" is a block of code that opens and closes with braces `{ }`
and may contain more configuration parameters. Below, `app "label-name" {}` is
the `app` stanza.

```hcl
app "label-name" {
  #...
}
```

A stanza is a parameter, but not all parameters are stanzas.

## Label

A label for a stanza is identified by the quotation marks `""`.
The `app` stanza takes a label named "label-name".

```hcl
app "label-name" {
  #...
}
```

## Entrypoint

The Waypoint Entrypoint is a special binary (named `waypoint-entrypoint`) that
launches your application and is used to enable a lot of
[Waypoint functionality](../docs/entrypoint#functionality).
The entrypoint is sometimes referred to as the "CEB" which stands for
"custom entrypoint binary."

## Client

A Waypoint client is anything that interacts with the server to inspect,
create, update, or destroy projects, applications, etc. The most common
Waypoint client is the CLI or web UI.

## Server

A Waypoint server is the running server that stores data and orchestrates
operations such as builds, deploys, etc. A client interacts with a server.
The server is deployed using `waypoint install`.
