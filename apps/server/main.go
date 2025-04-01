/*
Copyright © 2025 Michal Hodur @ itsopenmiso
*/
package main

import "github.com/itsopenmiso/openmiso/cmd"

func main() {
	err := cmd.ServerCmd.Execute()
	if err != nil {
		return
	}
}
