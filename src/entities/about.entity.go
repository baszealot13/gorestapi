package entities

type About struct {
	ID          uint   `json:"about_id"`
	Title       string `json:about_title`
	Description string `json:about_description`
}

func (about *About) TableName() string {
	return "abouts"
}
