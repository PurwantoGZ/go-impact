package models

//ImageFile ...
type ImageFile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Data []byte `json:"data"`
}
