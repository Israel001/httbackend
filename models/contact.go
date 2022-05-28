package models

type Contact struct {
	BaseModel
	Firstname string `gorm:"type:varchar(50) not null" json:"firstname"`
	Lastname  string `gorm:"type:varchar(50) not null" json:"lastname"`
	Email     string `gorm:"type:varchar(50) not null" json:"email"`
	Phone     string `gorm:"type:varchar(50) not null" json:"phone"`
	Message   string `gorm:"type:varchar(50) not null" json:"message"`
}
