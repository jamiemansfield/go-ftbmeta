package ftbmeta

type Version struct {
	ID        int    `json:"id"`
	Parent    int    `json:"parent"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Changelog string `json:"changelog"`
	Type      string `json:"type"`
	Updated   int64  `json:"updated"`

	Specs   *Specs    `json:"specs"`
	Targets []*Target `json:"targets"`
	Files   []*File   `json:"files"`

	// Additional
	Advisories []*Advisory `json:"advisories,omitempty"`
}

type Target struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Updated int64  `json:"updated"`
}

type File struct {
	ID         int    `json:"id"`
	Type       string `json:"type"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	Version    string `json:"version"`
	URL        string `json:"url"`
	Sha1       string `json:"sha1"`
	Size       int    `json:"size"`
	ClientOnly bool   `json:"clientonly"`
	ServerOnly bool   `json:"serveronly"`
	Optional   bool   `json:"optional"`
	Updated    int64  `json:"updated"`
}
