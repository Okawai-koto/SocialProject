package models

// user represents data about a record user.
type VerifyCodeTemplate struct {
	Email      string `json:"email"`
	VerifyCode string `json:"verifycode"`
}
