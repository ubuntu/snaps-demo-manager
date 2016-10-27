package state

import (
	"io/ioutil"
	"path"

	"github.com/ubuntu/snaps-manager/dirs"
	"github.com/ubuntu/snaps-manager/logger"

	"gopkg.in/yaml.v2"
)

type configFormat struct {
	Snaps SnapsProperty
}

// SnapsProperty maps all snaps on the system.
type SnapsProperty map[string]SnapProperty

// AllSnapsProperty for all snaps on the system.
var AllSnapsProperty SnapsProperty

// SnapProperty expose properties of one snap.
type SnapProperty struct {
	Enabled      bool
	Instructions []SnapInstruction
}

// SnapInstruction unit.
type SnapInstruction struct {
	Origin        string
	Wait          int
	Nextoperation string
}

func init() {
	f := path.Join(dirs.Data, "config.yaml")
	c := configFormat{}

	// load instructions
	d, err := ioutil.ReadFile(f)
	if err != nil {
		// no file available: can be just installed with no instructions
		return
	}
	if err = yaml.Unmarshal(d, &c); err != nil {
		logger.Err("Couldn't deserialized config from file:", err)
	}
	logger.Info("Loaded configuration: %+v\n", c)
	AllSnapsProperty = c.Snaps
}
