package models

type Company struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
