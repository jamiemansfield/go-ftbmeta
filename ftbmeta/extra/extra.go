package extra

import (
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta"
)

type PackExtras struct {
	Overrides *PackOverrides `json:"overrides"`

	Advisories   map[string][]*ftbmeta.Advisory `json:"advisories"`
	Links        map[string]string              `json:"links"`
	Servers      []*Server                      `json:"servers"`
	Availability *ftbmeta.Availability          `json:"availability"`
}

type PackOverrides struct {
	// The pack itself
	Synopsis    string         `json:"synopsis"`
	Description string         `json:"description"`
	ExtraTags   []*ftbmeta.Tag `json:"+tags"`

	// The pack's versions
	Changelogs map[string]string `json:"changelogs"`
}

type Server struct {
	Name     string `json:"name"`
	Website  string `json:"website"`
	Hostname string `json:"hostname"`
}
