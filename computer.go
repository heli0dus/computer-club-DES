package main

import (
	"fmt"
	"math"
)

type Computer struct {
	isOccupied bool
	user       string
	since      int //time in minutes
	cost       int
	revenue    int
	totaltime  int
}

func (c *Computer) Occupy(user string, time int) error {
	if c.isOccupied {
		return fmt.Errorf("PlaceIsBusy")
	}

	c.isOccupied = true
	c.user = user
	c.since = time
	return nil
}

func (c *Computer) Free(time int) error {
	if !c.isOccupied {
		return fmt.Errorf("PlaceNotOccupied")
	}

	c.isOccupied = false

	c.revenue += int(math.Ceil(float64(time-c.since)/60.0)) * c.cost
	c.totaltime += (time - c.since)
	return nil
}
