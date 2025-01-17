---
layout: docs
page_title: Introduction
description: |-
  Welcome to Waypoint! This introduction section covers what Waypoint is, the problem Waypoint aims to solve, and how Waypoint compares to other software.
---

# Introduction to Waypoint

Welcome to Waypoint! This introduction section covers what Waypoint is,
the problem Waypoint aims to solve, and how Waypoint compares to other software.
If you just want to dive into using Waypoint, head over to the
[getting started](../docs/getting-started) section.

## What is Waypoint?

Waypoint is a tool that enables developers to describe how to get their
applications from development to production in a single file and deploy
using a single command: `waypoint up`. After deployment, Waypoint provides
tools such as
[logs](../docs/logs),
[exec](../docs/exec),
and more to validate and debug any deployments. Waypoint is fully extensible
based on a [plugin system](../docs/plugins) which allows Waypoint to work
with any tools and platforms.

### Waypoint Goals

1. **Consistency of Workflows.** Waypoint aims to provide an easy-to-use, consistent
   workflow for getting applications from development to production: `waypoint up`. We believe
   that fragmentation in workflow based on platform is a big challenge for teams
   that attempt to use multiple build, deploy, release tooling.

2. **Confidence in Deployment.** The first thing any developer does after
   deploying is validate it works: open a browser, refresh the page, check logs,
   etc. Waypoint provides tools such as
   [logs](../docs/logs) and
   [exec](../docs/exec)
   to give you confidence that deployments succeed.

3. **Extensibility with the Ecosystem.** Consistency is only useful if
   Waypoint works with all the tools and patterns you work with. Waypoint is
   fully [extensible via a plugin system](../docs/plugins) that allows you to bring
   in custom builders, deployment platforms, and more.

## Why Waypoint?

Waypoint was built for one simple reason: developers just want to deploy.

But the modern day developer is being inundated with complexity: containers,
schedulers, YAML files, serverless, and more. This complexity has improved the
capability of our applications in many ways, but the cost can be seen in the
learning curve required to just get your first application deployed.
Developers just want to deploy.

Another challenge we saw is that depending on where you want to deploy
your application, the tool you use is often different. Docker and `kubectl`
for Kubernetes, Packer and Terraform for VMs, custom CLIs for each serverless
platform, etc. For the individual, this again poses a learning curve challenge.
For teams, this poses a challenge in consistency.

On one hand, **Waypoint was built to provide ease of use. You don't _need_ to
write Dockerfiles, YAML, etc. anymore.** We have plugins to automatically
detect your language, build an image, and deploy. Yes, you have to write
some minimal configuration but we're talking around 15 lines of text for one
tool versus hundreds of lines across different tools using different languages.

On the other hand, **Waypoint was built to be consistent across any platform.**
Waypoint can be extended with plugins to target any build, deploy, release logic.
For example, if you already have Dockerfiles, YAML, etc. written, we have plugins
that are able to utilize that. Or, if you have projects that use a different
paradigm such as VMs or Serverless, we have plugins for that.

Developers just want to deploy. Waypoint gets you there.

-> **Waypoint is an early project.** The sections above represent the background
for Waypoint and the goals we have for the project, but we know Waypoint
can continue to improve. We are iterating on Waypoint quickly to continue to make
it easier and more enjoyable to deploy applications.
