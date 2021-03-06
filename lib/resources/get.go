package resources

import (
	"fmt"
	"strings"

	"github.com/hofstadter-io/hof/cmd/hof/flags"
)

func RunGetFromArgs(args []string) error {
	labels := flags.RootPflags.Labels
	fmt.Println("lib/resources.Get")

	/*
	if flags.RootPflags.Global {
	} else if flags.RootPflags.Local {
	} else {
	}
	*/

	// TODO, handle config/secret/context like datamodel/modelset/model/view
	// (separate from builtin/custom resources, as they have different Cue entrypoints / file locations)

	// TODO, be lazy
	// load resources / datamodels into Cue runtime(s)

	// lookup things in the Cue values
	for _, arg := range args {
		resource := arg
		name := ""
		// resource/name ?
		if strings.Contains(arg, "/") {
			flds := strings.Split(arg, "/")
			if len(flds) != 2 {
				return fmt.Errorf("Resource should only have one or two parts: <resource>[/<name>]")
			}
			resource = flds[0]
			name = flds[1]
		}

		fmt.Println(" -", resource, name, labels)

		// check resource type, mayeb do different things
	}

	return nil
}
