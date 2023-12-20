package domain

type SihFile struct {
	UUID      string `json:"uuid"`
	FileName  string `json:"filename"`
	Timestamp string `json:"timestamp"`
	Size      int64  `json:"size"`
}
