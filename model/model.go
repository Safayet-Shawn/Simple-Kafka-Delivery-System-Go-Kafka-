package model

type Order struct {
	Id      int64    `json:"id"`
	Order   []string `json:"order"`
	Address string   `json:"address"`
}
