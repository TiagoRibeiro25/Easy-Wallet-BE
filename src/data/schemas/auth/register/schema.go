package schemas

type BodyData struct {
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	Password    string `json:"password"`
}

type ResponseData struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	DisplayName  string `json:"display_name"`
	Currency     string `json:"currency"`
	UserVerified bool   `json:"user_verified"`
	CreatedAt    string `json:"created_at"`
}
