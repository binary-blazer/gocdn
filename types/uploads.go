package types

type Upload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}