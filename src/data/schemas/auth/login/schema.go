package schemas

type BodyData struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type ResponseData struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	DisplayName  string `json:"display_name"`
	UserVerified bool   `json:"user_verified"`
	Currency     string `json:"currency"`
	CreatedAt    string `json:"created_at"`
}
