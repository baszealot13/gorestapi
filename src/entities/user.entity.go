package entities

type User struct {
	UserID         uint   `gorm:"primary_key" json:"user_id"`
	Username       string `gorm:"unique;not null;" json:"username"`
	Password       string `json:"password"`
	UserEmail      string `gorm:"unique;not null;" json:"user_email"`
	UserActive     bool   `gorm:"default:true" json:"user_active"`
	UserIsAdmin    bool   `gorm:"default:false" json:"user_is_admin"`
	UserCreated    string `gorm:"type:datetime"`
	UserCreatedBy  uint   `json:"user_created_by"`
	UserModified   string `gorm:"type:datetime"`
	UserModifiedBy uint   `json:"user_modified_by"`
}

func (user *User) TableName() string {
	return "users"
}
