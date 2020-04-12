package ftbmeta

type PackInfo struct {
	ID int `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Synopsis string `json:"synopsis"`
	Featured bool `json:"featured"`
	Type string `json:"type"`
	Updated int64 `json:"updated"`

	Icon *Art              `json:"icon"`
	Tags []*Tag `json:"tags"`
}
