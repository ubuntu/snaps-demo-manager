package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ubuntu/snaps-manager/client"
	"github.com/ubuntu/snaps-manager/manager"
	"github.com/ubuntu/snaps-manager/state"
)

var currentSnaps map[string]*manager.Snap

func main() {

	wg := sync.WaitGroup{}
	currentSnaps = make(map[string]*manager.Snap)

	for name, p := range state.AllSnapsProperty {
		startAndStopTracking(name, p, &wg)
	}
	wg.Done()
	return
	c := client.C
	snap, r, err := c.Snap("consul")
	fmt.Printf("%+v\n", snap)
	fmt.Printf("%+v\n", r)
	fmt.Printf("%+v\n", err)
	changeid, err := c.Install("face-detection-demo", &client.SnapOptions{
		Channel: "beta",
		DevMode: true,
	})
	fmt.Printf(changeid)
	fmt.Println(err)
	change, err := c.Change(changeid)
	fmt.Println(change.Status)
	for change.Status != "Done" {
		change, err = c.Change(changeid)
		fmt.Println("Waiting")
		fmt.Println(change.Status)
		fmt.Println(change.Ready)
		time.Sleep(time.Duration(time.Second))
	}
}

func startAndStopTracking(name string, p state.SnapProperty, wg *sync.WaitGroup) {
	s, trackingExists := currentSnaps[name]
	if p.Enabled && !trackingExists {
		s := manager.NewSnap(name)
		s.Track(wg)
		s.Instructions <- p.Instructions
		currentSnaps[name] = s
	} else if !p.Enabled && trackingExists {
		s.Stop <- true
		delete(currentSnaps, name)
	}
}
