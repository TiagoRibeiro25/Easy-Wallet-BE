package schemas

type ResponseData struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	IconID    string `json:"icon_id"`
	UserID    bool   `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
