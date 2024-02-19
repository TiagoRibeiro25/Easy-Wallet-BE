package schemas

type BodyData struct {
	Name        string `json:"name"`
	CategoryID  uint   `json:"category_id"`
	Cost        int    `json:"cost"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type ResponseData struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	CategoryID  uint   `json:"category_id"`
	Cost        int    `json:"cost"`
	Date        string `json:"date"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
