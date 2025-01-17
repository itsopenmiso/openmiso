---
layout: docs
page_title: 'TaskLauncher'
description: |-
  How to implement the TaskLauncher component for a Waypoint plugin
---

The `TaskLauncher` component makes possible launching Waypoint [On-Demand Runners](../docs/runner/on-demand-runner)
in a given platform, to perform [remote operations](../docs/projects/remote).

https://pkg.go.dev/github.com/hashicorp/waypoint-plugin-sdk/component#TaskLauncher

```go
type TaskLauncher interface {
	// StartTaskFunc should return a method for the "start task" operation.
	// This will have TaskLaunchInfo available to it to understand what the task
	// should do.
	StartTaskFunc() interface{}

	// StopTaskFunc is called to force a previously started task to stop. It will
	// be passed the state value returned by StartTaskFunc for identification.
	StopTaskFunc() interface{}

	// WatchTaskFunc is called after Start but before Stop to block and
	// watch a single task. It should stream output to the given UI and
	// return the exit status after it exits. It is given the state resulting
	// from StartTaskFunc so that it can look up the resource.
	WatchTaskFunc() interface{}
}
```

`TaskLauncher` has three functions which must be implemented, that all return
a function to be called by Waypoint. The function returned by `StartTaskFunc`
"starts" the on-demand runner in the platform. The function returned by
`StopTaskFunc` stops & removes the on-demand runner in the platform. Finally,
`WatchTaskFunc` returns a function that logs the stream of output from the
on-demand runner. This is useful for investigating a job by viewing its logs
with the [`waypoint job get-stream` command](../commands/job-get-stream).

```go
type TaskLauncher struct {
	config TaskLauncherConfig
}

type TaskLauncherConfig struct {
  // Other component fields
}

func (p *TaskLauncher) StartTaskFunc() interface{} {
	return p.StartTask
}

func (p *TaskLauncher) StopTaskFunc() interface{} {
	return p.StopTask
}

func (p *TaskLauncher) WatchTaskFunc() interface{} {
	return p.WatchTask
}

func (p *TaskLauncher) StartTask(
	ctx context.Context,
	log hclog.Logger,
	tli *component.TaskLaunchInfo,
) (*TaskInfo, error) {
  return nil, nil
}

func (p *TaskLauncher) StopTask(
	ctx context.Context,
	log hclog.Logger,
	ti *TaskInfo,
) error {
  return nil
}

func (p *TaskLauncher) WatchTask(
	ctx context.Context,
	log hclog.Logger,
	ui terminal.UI,
	ti *TaskInfo,
) (*component.TaskResult, error) {
  return nil, nil
}
```
