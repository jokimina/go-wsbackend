package model

type WasteItem struct {
	Name string `json:"name"`
	Qp   string `json:"qp"`
	FL   string `json:"fl"`
	Cats int64  `json:"cats"`
}
