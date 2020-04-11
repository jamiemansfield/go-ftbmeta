package ftbmeta

import (
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
)

type PackInfo struct {
	ID int `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Synopsis string `json:"synopsis"`
	Featured bool `json:"featured"`
	Type string `json:"type"`
	Updated int64 `json:"updated"`

	Icon *Art              `json:"icon"`
	Tags []*modpacksch.Tag `json:"tags"`
}
