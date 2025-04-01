package waypointfile

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"os"
	"path/filepath"
	"strings"
)

//FIXME: I'm ugly but when it's okay I get the job done...

func Synth() *Jobfile {
	var hclFiles []string

	dir, e := os.Getwd()
	if e != nil {
		panic(e)
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".hcl") {
			hclFiles = append(hclFiles, path) // Dodanie pliku do listy
		} else if !info.IsDir() && strings.HasSuffix(info.Name(), ".miso") {
			hclFiles = append(hclFiles, path) // Dodanie pliku do listy
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	parser := hclparse.NewParser()
	for _, f := range hclFiles {
		parser.ParseHCLFile(f)
	}

	jf := &Jobfile{}
	list := []*hcl.File{}
	for _, file := range parser.Files() {
		list = append(list, file)
	}
	gohcl.DecodeBody(hcl.MergeFiles(list), nil, jf)
	return jf

}
