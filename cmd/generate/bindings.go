package main

import (
	"github.com/gosimple/slug"
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta"
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta/extra"
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
	"sort"
	"strings"
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

func convertAuthors(authors []*modpacksch.Author) []*ftbmeta.Author {
	var newAuthors []*ftbmeta.Author
	for _, author := range authors {
		newAuthors = append(newAuthors, convertAuthor(author))
	}
	return newAuthors
}

func convertAuthor(author *modpacksch.Author) *ftbmeta.Author {
	return &ftbmeta.Author{
		ID:      author.ID,
		Name:    author.Name,
		Type:    author.Type,
		Website: author.Website,
		Updated: author.Updated,
	}
}

func convertTags(tags []*modpacksch.Tag) []*ftbmeta.Tag {
	var newTags []*ftbmeta.Tag
	for _, author := range tags {
		newTags = append(newTags, convertTag(author))
	}
	return newTags
}

func convertTag(tag *modpacksch.Tag) *ftbmeta.Tag {
	return &ftbmeta.Tag{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

// full

func convertPack(pack *modpacksch.Pack, extras *extra.PackExtras) *ftbmeta.Pack {
	// Handle overrides
	synopsis := pack.Synopsis
	if extras.Overrides.Synopsis != "" {
		synopsis = extras.Overrides.Synopsis
	}

	// Sort pack versions in increasing order
	sort.Sort(versionsByLatest(pack.Versions))

	// Create latest map
	latest := map[string]string{}
	for _, version := range pack.Versions {
		latest[strings.ToLower(version.Type)] = slug.MakeLang(version.Name, "en")
	}

	return &ftbmeta.Pack{
		ID:          pack.ID,
		Slug:        slug.MakeLang(pack.Name, "en"),
		Name:        pack.Name,
		Synopsis:    synopsis,
		Description: pack.Description,
		Featured:    pack.Featured,
		Type:        strings.ToLower(pack.Type),
		Updated:     getPackLastUpdated(pack),
		Art:         convertArtMap(pack.Art),
		Authors:     convertAuthors(pack.Authors),
		Versions:    convertVersionInfos(pack.Versions),
		Tags:        convertTags(pack.Tags),
		Links:       extras.Links,
		Latest:      latest,
		Advisories:  extras.Advisories["current"],
	}
}

func convertVersionInfo(version *modpacksch.VersionInfo) *ftbmeta.VersionInfo {
	return &ftbmeta.VersionInfo{
		ID:      version.ID,
		Slug:    slug.MakeLang(version.Name, "en"),
		Name:    version.Name,
		Type:    strings.ToLower(version.Type),
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
		Type:     strings.ToLower(pack.Type),
		Updated:  getPackLastUpdated(pack),
		Icon:     convertArt(pack.GetIcon()),
		Tags:     convertTags(pack.Tags),
	}
}

// version

func convertTargets(targets []*modpacksch.Target) []*ftbmeta.Target {
	var newTargets []*ftbmeta.Target
	for _, target := range targets {
		newTargets = append(newTargets, convertTarget(target))
	}
	return newTargets
}

func convertTarget(target *modpacksch.Target) *ftbmeta.Target {
	return &ftbmeta.Target{
		ID:      target.ID,
		Type:    target.Type,
		Name:    target.Name,
		Version: target.Version,
		Updated: target.Updated,
	}
}

func convertFiles(files []*modpacksch.File) []*ftbmeta.File {
	var newFiles []*ftbmeta.File
	for _, file := range files {
		newFiles = append(newFiles, convertFile(file))
	}
	return newFiles
}

func convertFile(file *modpacksch.File) *ftbmeta.File {
	return &ftbmeta.File{
		ID:         file.ID,
		Type:       file.Type,
		Path:       file.Path,
		Name:       file.Name,
		Version:    file.Version,
		URL:        file.URL,
		Sha1:       file.Sha1,
		Size:       file.Size,
		ClientOnly: file.ClientOnly,
		ServerOnly: file.ServerOnly,
		Optional:   file.Optional,
		Updated:    file.Updated,
	}
}

func convertVersion(version *modpacksch.Version, changelog *modpacksch.VersionChangelog, extras *extra.PackExtras) *ftbmeta.Version {
	versionSlug := slug.MakeLang(version.Name, "en")

	changelogRaw := changelog.Content
	if extras.Changelogs[versionSlug] != "" {
		changelogRaw = extras.Changelogs[versionSlug]
	}

	return &ftbmeta.Version{
		ID:         version.ID,
		Parent:     version.Parent,
		Slug:       versionSlug,
		Name:       version.Name,
		Changelog:  changelogRaw,
		Type:       strings.ToLower(version.Type),
		Updated:    version.Updated,
		Specs:      convertSpecs(version.Specs),
		Targets:    convertTargets(version.Targets),
		Files:      convertFiles(version.Files),
		Advisories: extras.Advisories[versionSlug],
	}
}
