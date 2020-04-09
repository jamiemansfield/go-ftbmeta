package ftbmeta

import (
	"github.com/gosimple/slug"
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta/util"
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
)

type Pack struct {
	ID int `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Synopsis string `json:"synopsis"`
	Description string `json:"description"`
	Featured bool `json:"featured"`
	Type string `json:"type"`
	Updated int64 `json:"updated"`

	Art map[string]*Art          `json:"art"`
	Authors []*modpacksch.Author `json:"authors"`
	Versions []*VersionInfo      `json:"versions"`
	Tags []*modpacksch.Tag       `json:"tags"`
}

func NewPack(pack *modpacksch.Pack) *Pack {
	return &Pack{
		ID:          pack.ID,
		Slug:        slug.MakeLang(pack.Name, "en"),
		Name:        pack.Name,
		Synopsis:    pack.Synopsis,
		Description: pack.Description,
		Featured:    pack.Featured,
		Type:        pack.Type,
		Updated:     util.GetPackLastUpdated(pack),
		Art:         ConvertArt(pack.Art),
		Authors:     pack.Authors,
		Versions:    ConvertVersionInfos(pack.Versions),
		Tags:        pack.Tags,
	}
}

func ConvertArt(art []*modpacksch.Art) map[string]*Art {
	var newArt = map[string]*Art{}
	for _, piece := range art {
		newArt[piece.Type] = NewArt(piece)
	}
	return newArt
}

type VersionInfo struct {
	ID int `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Type string `json:"type"`
	Updated int64 `json:"updated"`
	Specs *Specs `json:"specs"`
}

func NewVersionInfo(version *modpacksch.VersionInfo) *VersionInfo {
	return &VersionInfo{
		ID:      version.ID,
		Slug:    slug.MakeLang(version.Name, "en"),
		Name:    version.Name,
		Type:    version.Type,
		Updated: version.Updated,
		Specs:   NewSpecs(version.Specs),
	}
}

func ConvertVersionInfos(versions []*modpacksch.VersionInfo) []*VersionInfo {
	var infos []*VersionInfo
	for _, version := range versions {
		infos = append(infos, NewVersionInfo(version))
	}
	return infos
}

type Specs struct {
	Minimum int `json:"minimum"`
	Recommended int `json:"recommended"`
}

func NewSpecs(specs *modpacksch.Specs) *Specs {
	return &Specs{
		Minimum:     specs.Minimum,
		Recommended: specs.Recommended,
	}
}
