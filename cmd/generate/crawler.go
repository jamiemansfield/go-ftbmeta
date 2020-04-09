package main

import "github.com/jamiemansfield/go-modpacksch/modpacksch"

type crawler struct {
	Terms []string
}

func newCrawler() *crawler {
	return &crawler{
		Terms: []string{},
	}
}

func (c *crawler) useTerm(term... string) *crawler {
	c.Terms = append(c.Terms, term...)
	return c
}

func (c *crawler) findPacks(client *modpacksch.Client) ([]*modpacksch.Pack, error) {
	var ids []int
	for _, term := range c.Terms {
		packs, err := client.Packs.Search(50, term)
		if err != nil {
			return nil, err
		}

		ids = append(ids, packs...)
	}

	var full []*modpacksch.Pack
	for _, id := range ids {
		pack, err := client.Packs.GetPack(id)
		if err != nil {
			return nil, err
		}

		full = append(full, pack)
	}
	return full, nil
}
