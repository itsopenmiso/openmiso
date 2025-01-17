---
layout: docs
page_title: path - Variables - waypoint.hcl
description: |-
  Path variables are used to construct an absolute path to reference files or directories that may be used within your Waypoint configuration.
---

# `path` Variables

Path variables are used to construct an absolute path to reference
files or directories that may be used within your Waypoint configuration.
For example, path variables may be used with the Docker builder to
specify the path to the Dockerfile to build.

## `path.app`

<Placement groups={[['app']]} />

The directory of the application being built or deployed with Waypoint.

If [`path`](../docs/waypoint-hcl/app#path) is not set for the app, this will
always equal `path.project`. If you are referencing app-relative data,
you should always use `path.app` in case you move the application in
the future.

Example:

```hcl
dockerfile = "${path.app}/Dockerfile"
```

## `path.project`

<Placement groups={[['* (global)']]} />

The directory that contains the `waypoint.hcl`. **This is the recommended
variable to use** when referencing files relative to the `waypoint.hcl` file.

Unlike `path.pwd`, this will always point to the directory where the
`waypoint.hcl` is. If you execute Waypoint from a subdirectory and it loads
a `waypoint.hcl` file in some parent directory, that directory will be
the value of `path.project`.
When executing the `waypoint` CLI in the same directory as the `waypoint.hcl`
file, `path.pwd` and `path.project` are equivalent.

Example:

```hcl
dockerfile = "${path.project}/Dockerfile"
```

## `path.pwd`

<Placement groups={[['* (global)']]} />

The working directory for the CLI. This is equivalent to the result of
`pwd` in the shell.

When executed from a subdirectory, Waypoint will search parent directories
for a `waypoint.hcl`. In this case, the `path.pwd` variable will point to
the subdirectory where you executed Waypoint, _not_ the directory where
the `waypoint.hcl` is.

Example:

```hcl
dockerfile = "${path.pwd}/Dockerfile"
```

~> **`path.pwd` is not recommended.** You almost always want `path.project`
or `path.app` instead. Using `path.pwd` makes Waypoint execution dependent
on the current working directory when executing `waypoint`, which can be
surprising to users.
