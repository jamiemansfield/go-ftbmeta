package main

import (
	"github.com/gosimple/slug"
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta"
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta/extra"
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
)

// common

func convertArt(art *modpacksch.Art) *ftbmeta.Art {
	return &ftbmeta.Art{
		URL:     art.URL,
		Width:   art.Width,
		Height:  art.Height,
		Sha1:    art.Sha1,
		Size:    art.Size,
		Updated: art.Updated,
	}
}

func convertArtMap(art []*modpacksch.Art) map[string]*ftbmeta.Art {
	var newArt = map[string]*ftbmeta.Art{}
	for _, piece := range art {
		newArt[piece.Type] = convertArt(piece)
	}
	return newArt
}

func convertSpecs(specs *modpacksch.Specs) *ftbmeta.Specs {
	return &ftbmeta.Specs{
		Minimum:     specs.Minimum,
		Recommended: specs.Recommended,
	}
}

// full

func convertPack(pack *modpacksch.Pack, extras *extra.PackExtras) *ftbmeta.Pack {
	synopsis := pack.Synopsis
	if extras.Overrides.Synopsis != "" {
		synopsis = extras.Overrides.Synopsis
	}

	return &ftbmeta.Pack{
		ID:          pack.ID,
		Slug:        slug.MakeLang(pack.Name, "en"),
		Name:        pack.Name,
		Synopsis:    synopsis,
		Description: pack.Description,
		Featured:    pack.Featured,
		Type:        pack.Type,
		Updated:     getPackLastUpdated(pack),
		Art:         convertArtMap(pack.Art),
		Authors:     pack.Authors,
		Versions:    convertVersionInfos(pack.Versions),
		Tags:        pack.Tags,
		Links:       extras.Links,
	}
}

func convertVersionInfo(version *modpacksch.VersionInfo) *ftbmeta.VersionInfo {
	return &ftbmeta.VersionInfo{
		ID:      version.ID,
		Slug:    slug.MakeLang(version.Name, "en"),
		Name:    version.Name,
		Type:    version.Type,
		Updated: version.Updated,
		Specs:   convertSpecs(version.Specs),
	}
}

func convertVersionInfos(versions []*modpacksch.VersionInfo) []*ftbmeta.VersionInfo {
	var infos []*ftbmeta.VersionInfo
	for _, version := range versions {
		infos = append(infos, convertVersionInfo(version))
	}
	return infos
}

// simple

func convertPackInfo(pack *modpacksch.Pack) *ftbmeta.PackInfo {
	return &ftbmeta.PackInfo{
		ID:       pack.ID,
		Slug:     slug.MakeLang(pack.Name, "en"),
		Name:     pack.Name,
		Synopsis: pack.Synopsis,
		Featured: pack.Featured,
		Type:     pack.Type,
		Updated:  getPackLastUpdated(pack),
		Icon:     convertArt(pack.GetIcon()),
		Tags:     pack.Tags,
	}
}

// version

func convertVersion(version *modpacksch.Version, changelog *modpacksch.VersionChangelog) *ftbmeta.Version {
	return &ftbmeta.Version{
		ID:        version.ID,
		Parent:    version.Parent,
		Slug:      slug.MakeLang(version.Name, "en"),
		Name:      version.Name,
		Changelog: changelog.Content,
		Type:      version.Type,
		Updated:   version.Updated,
		Specs:     convertSpecs(version.Specs),
		Targets:   version.Targets,
		Files:     version.Files,
	}
}
