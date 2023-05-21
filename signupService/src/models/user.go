package models

// user represents data about a record user.
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
