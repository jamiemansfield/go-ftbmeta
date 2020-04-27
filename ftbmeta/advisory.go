package ftbmeta

type Advisory struct {
	Severity string `json:"severity"`
	Message string `json:"message"`
	URL string `json:"url,omitempty"`
}
