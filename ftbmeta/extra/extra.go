package extra

import (
	"encoding/json"
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta"
	"io/ioutil"
)

type PackExtras struct {
	Overrides struct {
		Synopsis   string            `json:"synopsis"`
		ExtraTags  []*ftbmeta.Tag    `json:"+tags"`
		Changelogs map[string]string `json:"changelogs"`
	} `json:"overrides"`

	Advisories map[string][]*ftbmeta.Advisory `json:"advisories"`
	Links      map[string]string              `json:"links"`
	Servers    []*Server                      `json:"servers"`
}

func GetPackExtras(root string, pack string) (*PackExtras, error) {
	path := root + "/pack/" + pack + ".json"

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var extra PackExtras
	err = json.Unmarshal(bytes, &extra)
	if err != nil {
		return nil, err
	}
	return &extra, nil
}

type Server struct {
	Name     string `json:"name"`
	Website  string `json:"website"`
	Hostname string `json:"hostname"`
}
