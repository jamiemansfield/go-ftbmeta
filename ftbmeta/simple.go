package ftbmeta

import (
	"github.com/gosimple/slug"
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta/util"
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

func NewPackInfo(pack *modpacksch.Pack) *PackInfo {
	return &PackInfo{
		ID:       pack.ID,
		Slug:     slug.MakeLang(pack.Name, "en"),
		Name:     pack.Name,
		Synopsis: pack.Synopsis,
		Featured: pack.Featured,
		Type:     pack.Type,
		Updated:  util.GetPackLastUpdated(pack),
		Icon:     NewArt(pack.GetIcon()),
		Tags:     pack.Tags,
	}
}

type Art struct {
	URL string `json:"url"`
	Width int `json:"width"`
	Height int `json:"height"`
	Sha1 string `json:"sha1"`
	Size int `json:"size"`
	Updated int64 `json:"updated"`
}

func NewArt(art *modpacksch.Art) *Art {
	return &Art{
		URL:     art.URL,
		Width:   art.Width,
		Height:  art.Height,
		Sha1:    art.Sha1,
		Size:    art.Size,
		Updated: art.Updated,
	}
}
