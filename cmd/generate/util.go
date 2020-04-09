package main

import (
	"encoding/json"
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
	"io/ioutil"
)

// GENERATION
// ----------

func writeJson(destination string, v interface{}) error {
	out, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(destination, out, 0644)
}

// SORTING
// -------

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
