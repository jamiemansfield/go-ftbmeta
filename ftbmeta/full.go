package ftbmeta

type Pack struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Synopsis    string `json:"synopsis"`
	Description string `json:"description"`
	Featured    bool   `json:"featured"`
	Type        string `json:"type"`
	Updated     int64  `json:"updated"`

	Art      map[string]*Art `json:"art"`
	Authors  []*Author       `json:"authors"`
	Versions []*VersionInfo  `json:"versions"`
	Tags     []*Tag          `json:"tags"`

	// Additional
	Links  map[string]string `json:"links,omitempty"`
	Latest map[string]string `json:"latest"`
}

type VersionInfo struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Updated int64  `json:"updated"`
	Specs   *Specs `json:"specs"`
}
