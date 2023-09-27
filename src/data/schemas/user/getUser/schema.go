package schemas

type ResponseData struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	DisplayName  string `json:"display_name"`
	UserVerified bool   `json:"user_verified"`
	Currency     string `json:"currency"`
	CreatedAt    string `json:"created_at"`
}
