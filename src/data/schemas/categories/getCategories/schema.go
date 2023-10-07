package schemas

type ResponseData struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	IconID    uint   `json:"icon_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
