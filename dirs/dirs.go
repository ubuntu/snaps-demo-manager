package dirs

import (
	"log"
	"os"
	"path/filepath"
)

// Data access to write storage path
var Data string

func init() {
	// Set main set of directories

	Data = os.Getenv("SNAP_DATA")
	if Data == "" {
		root, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		Data = root
	}
}
