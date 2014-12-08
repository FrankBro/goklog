// Copyright (c) 2014 Datacratic. All rights reserved.

package klog

import (
	"fmt"
	"time"
)

type Line struct {
	Timestamp time.Time `json:"ts"`
	Key       string    `json:"key"`
	Value     string    `json:"val"`
}

func (line *Line) String() string {
	return fmt.Sprintf("%s <%s> %s", line.Timestamp, line.Key, line.Value)
}

type lineArray []*Line

func (array lineArray) Len() int           { return len(array) }
func (array lineArray) Swap(i, j int)      { array[i], array[j] = array[j], array[i] }
func (array lineArray) Less(i, j int) bool { return array[i].Timestamp.Before(array[j].Timestamp) }
