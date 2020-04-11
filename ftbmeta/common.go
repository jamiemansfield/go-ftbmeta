package ftbmeta

type Art struct {
	URL string `json:"url"`
	Width int `json:"width"`
	Height int `json:"height"`
	Sha1 string `json:"sha1"`
	Size int `json:"size"`
	Updated int64 `json:"updated"`
}

type Specs struct {
	Minimum int `json:"minimum"`
	Recommended int `json:"recommended"`
}
