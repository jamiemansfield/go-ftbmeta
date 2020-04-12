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

type Author struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Website string `json:"website"`
	Updated int64 `json:"updated"`
}

type Tag struct {
	ID int `json:"id"`
	Name string `json:"name"`
}
