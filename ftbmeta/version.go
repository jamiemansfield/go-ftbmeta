package ftbmeta

import (
	"github.com/gosimple/slug"
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

func NewVersion(version *modpacksch.Version, changelog *modpacksch.VersionChangelog) *Version {
	return &Version{
		ID:        version.ID,
		Parent:    version.Parent,
		Slug:      slug.MakeLang(version.Name, "en"),
		Name:      version.Name,
		Changelog: changelog.Content,
		Type:      version.Type,
		Updated:   version.Updated,
		Specs:     NewSpecs(version.Specs),
		Targets:   version.Targets,
		Files:     version.Files,
	}
}
