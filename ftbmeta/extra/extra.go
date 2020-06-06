package extra

import (
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta"
)

type PackExtras struct {
	Overrides *PackOverrides `json:"overrides"`

	Advisories map[string][]*ftbmeta.Advisory `json:"advisories"`
	Links      map[string]string              `json:"links"`
	Servers    []*Server                      `json:"servers"`
}

type PackOverrides struct {
	Synopsis   string            `json:"synopsis"`
	ExtraTags  []*ftbmeta.Tag    `json:"+tags"`
	Changelogs map[string]string `json:"changelogs"`
}

type Server struct {
	Name     string `json:"name"`
	Website  string `json:"website"`
	Hostname string `json:"hostname"`
}
