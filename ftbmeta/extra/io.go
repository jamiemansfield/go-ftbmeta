package extra

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/jamiemansfield/go-ftbmeta/ftbmeta"
)

func GetPackExtras(root string, pack string) (*PackExtras, error) {
	path := filepath.Join(root, "pack", pack + ".json")

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

func SafeGetPackExtras(root string, pack string) *PackExtras {
	extras, err := GetPackExtras(root, pack)
	if err != nil {
		// Create a PackExtras with none-nil maps
		extras = &PackExtras{
			Overrides:  &PackOverrides{
				Changelogs: map[string]string{},
			},
			Advisories: map[string][]*ftbmeta.Advisory{},
			Links: map[string]string{},
		}
	}

	return extras
}
