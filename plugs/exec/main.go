package main

import "github.com/itsopenmiso/openmiso/pkg/plugsdk/misoplug"

type Exec struct {
	misoplug.MisoPlug
}

func main() {
	s := &Exec{}
	s.WithBuilder(s).
		WithReleaser(s).
		WithDeployer(s).
		Run()
}
