package ftbmeta

import (
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
)

type Version struct {
	ID        int    `json:"id"`
	Parent    int    `json:"parent"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Changelog string `json:"changelog"`
	Type      string `json:"type"`
	Updated   int64  `json:"updated"`

	Specs   *Specs               `json:"specs"`
	Targets []*modpacksch.Target `json:"targets"`
	Files   []*modpacksch.File   `json:"files"`
}
