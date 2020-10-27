package model

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
}
