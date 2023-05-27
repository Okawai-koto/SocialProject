package models

type MailData struct {
	Email      string `json:"email"`
	Verifycode string `json:"verifycode"`
}
