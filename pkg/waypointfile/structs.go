package waypointfile

import "github.com/hashicorp/hcl/v2"

type Jobfile struct {
	Project   string      `hcl:"project,optional"`
	Apps      []*App      `hcl:"app,block"`
	Variables []*Variable `hcl:"variable,block"`
	Secrets   []*Variable `hcl:"secret,block"`
}

type App struct {
	Name          string          `hcl:"name,label"`
	BuildBlocks   []*BuildBlock   `hcl:"build,block"`
	DeployBlocks  []*DeployBlock  `hcl:"deploy,block"`
	ReleaseBlocks []*ReleaseBlock `hcl:"release,block"`
}

type BuildBlock struct {
	Use        *UseBlock `hcl:"use,block"`
	StepConfig hcl.Body  `hcl:",remain"`
}

type DeployBlock struct {
	Use        *UseBlock `hcl:"use,block"`
	StepConfig hcl.Body  `hcl:",remain"`
}

type ReleaseBlock struct {
	Use        *UseBlock `hcl:"use,block"`
	StepConfig hcl.Body  `hcl:",remain"`
}

type UseBlock struct {
	PluginName string `hcl:"name,label"`
}

type Variable struct {
	Name      string   `hcl:"name"`
	Value     string   `hcl:"value"`
	Remaining hcl.Body `hcl:",remain"`
}
