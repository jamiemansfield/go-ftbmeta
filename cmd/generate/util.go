package main

import (
	"encoding/json"
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
	"io/ioutil"
)

// GENERATION
// ----------

func writeJson(destination string, v interface{}) error {
	out, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(destination, out, 0644)
}

// SORTING
// -------

// Arrange packs by their ID in increasing order.
type packsById []*modpacksch.Pack

func (p packsById) Len() int {
	return len(p)
}

func (p packsById) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p packsById) Less(i, j int) bool {
	return p[i].ID < p[j].ID
}

// Arrange packs by their updated time in increasing order.
type versionsByLatest []*modpacksch.VersionInfo

func (p versionsByLatest) Len() int {
	return len(p)
}

func (p versionsByLatest) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p versionsByLatest) Less(i, j int) bool {
	return p[i].Updated < p[j].Updated
}

