package main

import (
	"fmt"
	"time"

	"github.com/ubuntu/snaps-manager/client"
)

func main() {
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
