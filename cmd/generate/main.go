package main

import (
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
}
