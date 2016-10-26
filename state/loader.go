package state

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/ubuntu/snaps-manager/dirs"

	"gopkg.in/yaml.v2"
)

type instructionsFormat struct {
	Snaps map[string][]struct {
		Origin string
		Wait   int
		Next   string
	}
}

// SnapsInstructions are instructions global to the system
var SnapsInstructions instructionsFormat

func init() {
	f := path.Join(dirs.Data, "instructions.yaml")

	// load instructions
	d, err := ioutil.ReadFile(f)
	if err != nil {
		// no file available: can be just installed with no instructions
		return
	}
	if err = yaml.Unmarshal(d, &SnapsInstructions); err != nil {
		panic("Couldn't deserialized instructions from file:" + err.Error())
	}

	fmt.Printf("%+v\n", SnapsInstructions)

}

func F() {

}
