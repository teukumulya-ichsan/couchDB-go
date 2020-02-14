package models

type User struct {
	ID      string `json:"_id"`
	Rev     string `json:"_rev,omitempty"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}