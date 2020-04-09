package main

import (
	"fmt"
	"github.com/jamiemansfield/go-ftbmeta/ftbmeta"
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
	"os"
	"path/filepath"
	"sort"
)

var (
	DEST = "_site"
)

func main() {
	// Setup modpacks.ch crawler, with reasonable defaults
	client := modpacksch.NewClient(nil)
	packs, err := newCrawler().useTerm(
		"1.15",
		"1.14",
		"1.13",
		"1.12",
		"1.11",
		"1.10",
		"1.7.10",
	).findPacks(client)
	if err != nil {
		panic(err)
	}
	sort.Sort(packsById(packs))

	// ROUTE: /packs/
	packsPath := filepath.Join(DEST, "packs")
	if err = os.MkdirAll(packsPath, os.ModePerm); err != nil {
		panic(err)
	}

	// -> Create ftbmeta.PackInfo's
	var packInfos []*ftbmeta.PackInfo
	for _, pack := range packs {
		packInfos = append(packInfos, ftbmeta.NewPackInfo(pack))
	}

	// -> Write to /packs/index.json
	err = writeJson(filepath.Join(packsPath, "index.json"), packInfos)
	if err != nil {
		panic(err)
	}

	// ROUTE: /pack/{slug}/ AND /pack/{pack slug}/{version slug}/
	packPath := filepath.Join(DEST, "pack")
	if err = os.MkdirAll(packPath, os.ModePerm); err != nil {
		panic(err)
	}

	// -> Create pack entries
	for _, pack := range packs {
		fullPack := ftbmeta.NewPack(pack)

		path := filepath.Join(packPath, fullPack.Slug)
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			panic(err)
		}

		// -> Write to /pack/{slug}/index.json
		err = writeJson(filepath.Join(path, "index.json"), fullPack)
		if err != nil {
			panic(err)
		}

		for _, versionInfo := range pack.Versions {
			version, err := client.Packs.GetVersion(pack.ID, versionInfo.ID)
			if err != nil {
				panic(err)
			}
			changelog, err := client.Packs.GetVersionChangelog(pack.ID, versionInfo.ID)
			if err != nil {
				fmt.Printf("%s %s missing changelog: %s\n", pack.Name, version.Name, err)
				changelog = &modpacksch.VersionChangelog{}
			}

			fullVersion := ftbmeta.NewVersion(version, changelog)

			versionPath := filepath.Join(path, fullVersion.Slug)
			if err = os.MkdirAll(versionPath, os.ModePerm); err != nil {
				panic(err)
			}

			// -> Write to /pack/{pack slug}/{version slug}/index.json
			err = writeJson(filepath.Join(versionPath, "index.json"), fullVersion)
			if err != nil {
				panic(err)
			}
		}
	}
}
