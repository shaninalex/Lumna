package user

type profileDTO struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Active   bool   `json:"active"`
}

type changePasswordDTO struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}
