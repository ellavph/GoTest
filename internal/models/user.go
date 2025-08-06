package models

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"-"`
	CompanyID *int64   `json:"company_id"`
	Company   *Company `json:"company,omitempty"`
}
