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
	Advisories   []*Advisory       `json:"advisories,omitempty"`
	Links        map[string]string `json:"links,omitempty"`
	Latest       map[string]string `json:"latest"`
	Availability *Availability     `json:"availability,omitempty"`
}

type VersionInfo struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Updated int64  `json:"updated"`
	Specs   *Specs `json:"specs"`
}

// Availability exposes the continual availability of a pack on an
// alternate platform.
type Availability struct {
	Curse struct {
		ID  int    `json:"id"`
		URL string `json:"url"`
	} `json:"curse"`
}
