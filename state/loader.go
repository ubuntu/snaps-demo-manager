package state

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/ubuntu/snaps-manager/dirs"

	"gopkg.in/yaml.v2"
)

type configFormat struct {
	Snaps snapsInstructions
}

type snapsInstructions map[string][]struct {
	Origin string
	Wait   int
	Next   string
}

// SnapsInstructions are instructions global to the system
var SnapsInstructions snapsInstructions

func init() {
	f := path.Join(dirs.Data, "config.yaml")
	var c configFormat
	// load instructions
	d, err := ioutil.ReadFile(f)
	if err != nil {
		// no file available: can be just installed with no instructions
		return
	}
	if err = yaml.Unmarshal(d, &c); err != nil {
		panic("Couldn't deserialized config from file:" + err.Error())
	}
	SnapsInstructions = c.Snaps
	fmt.Printf("%+v\n", SnapsInstructions)

}

func F() {

}
