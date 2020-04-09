package util

import (
	"github.com/jamiemansfield/go-modpacksch/modpacksch"
	"sort"
)

type newestDate []int64

func (p newestDate) Len() int {
	return len(p)
}

func (p newestDate) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p newestDate) Less(i, j int) bool {
	return p[i] > p[j]
}

func GetPackLastUpdated(pack *modpacksch.Pack) int64 {
	var dates []int64
	dates = append(dates, pack.Updated)
	for _, version := range pack.Versions {
		dates = append(dates, version.Updated)
	}
	sort.Sort(newestDate(dates))
	return dates[0]
}
