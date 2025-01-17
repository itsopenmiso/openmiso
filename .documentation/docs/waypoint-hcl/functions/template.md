---
layout: docs
page_title: Templating - Functions - waypoint.hcl
description: |-
  Waypoint has a set of functions for templating strings, files, and directories with both predefined and custom variable values. These templates can use conditionals, loops, and more to dynamically generate content.
---

# Templating Functions

Waypoint has a set of functions for templating strings, files, and directories
with both predefined and custom variable values. These templates can use
conditionals, loops, and more to dynamically generate content.

Waypoint has three templating functions documented below: `templatefile`,
`templatedir`, and `templatestring`. All of these functions use the same
templating syntax. The templates have access to all the same variables
and functions that are available at the function call-site (where the
function such as `templatedir` is called). Custom variables can also
be provided to the template.

## Syntax

Template syntax uses HCL [string templates](../docs/waypoint-hcl/syntax/expressions#string-templates).
This supports conditionals, loops, function calls, etc.

## `templatefile`

```hcl
templatefile(path[, vars...]) path
```

`templatefile` takes a path to a file and optionally one or more maps of variables
for the template, renders the template at `path`, and returns the path to
the rendered template. The returned path is a copy of the original template
in a temporary location. This function _does not modify the original file_.

**Reminder:** This function returns a file path, not the rendered contents of the file.
If you want the contents of the rendered file, you can wrap this function
with [`file`](../docs/waypoint-hcl/functions/all#file).

This function is useful for parameters that expect file paths. The example
below shows using this function with Docker to template the Dockerfile:

```hcl
app "web" {
  build {
    use "docker" {
      dockerfile = templatefile("${path.app}/Dockerfile.tpl", {
        theme = "rainbow"
      })
    }
  }
}
```

## `templatedir`

```hcl
templatedir(path[, vars...]) path
```

`templatedir` is similar to `templatefile`, except it takes a path to
a directory, treats all files within that directory as a template, copies
the results to a new temporary directory, and returns the temporary
directory path.

This function will recursively render all files in `path`.

This function is useful for parameters that expect directory paths, such
as specifying a directory of YAML files for Kubernetes.

## `templatestring`

```hcl
templatestring(string[, vars...]) string
```

`templatestring` renders a string as a template directly and returns
the rendered contents.

-> **If you have a static template string, you probably want to use
string literals instead.** This function is primarily for dynamically
sourced templates (such as reading a file or the contents of some
other variable). If you have a static string you can use HCL interpolations
directly, for example: `"Port: ${deploy.port}"`

Example, rendering a template file and returning the contents directly:

```hcl
userdata = templatestring(file("${path.app}/userdata.tpl"))
```
