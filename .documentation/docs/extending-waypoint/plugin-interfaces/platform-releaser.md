---
layout: docs
page_title: 'PlatformReleaser'
description: |-
  How to implement the PlatformReleaser component for a Waypoint plugin
---

# PlatformReleaser

https://pkg.go.dev/github.com/hashicorp/waypoint-plugin-sdk/component#PlatformReleaser

The `PlatformReleaser` interface is an optional interface that a Platform component
can implement to provide default [`Release`](../docs/extending-waypoint/plugin-interfaces/release-manager) functionality. This only takes effect
if no release is configured, therefore users cannot set custom configurations
for default releasers. Releases will happen in a way that is wholly determined
by the output of the deployment.

For example, if a deployment creates a target group in AWS, a default releaser
may manipulate that target group to route 100% of incoming traffic to the new
target group. The user need not set an identifier for the target group, because
the releaser will already have the necessary information from the deployment to
"release" it.

```go
type PlatformReleaser interface {
  // DefaultReleaserFunc() should return a function that returns
  // a ReleaseManger implementation. This ReleaseManager will NOT have
  // any config so it must work by default.
  DefaultReleaserFunc() interface{}
}
```

`PlatformReleaser` has a single function which must be implemented which
returns a function called by Waypoint to release the deployment. The function
returns the Releaser struct to Waypoint, which executes the Releaser's
`Release` function, as a `ReleaseManager` would.

```go
type Platform struct {
  // Other component fields
}

type Releaser struct {
  p *Platform
}

func (p *Platform) DefaultReleaserFunc() interface{} {
	return func() *Releaser { return &Releaser{p: p} }
}

func (r *Releaser) Release(
	ctx context.Context,
	log hclog.Logger,
	src *component.Source,
	ui terminal.UI,
	target *Deployment,
) (*Release, error)
```
