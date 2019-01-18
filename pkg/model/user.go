package model

type User struct {
	ID       string `json:"id,omitempty"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []*User
