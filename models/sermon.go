package models

type Sermon struct {
	BaseModel
	Title    string `gorm:"type:varchar(50) not null" json:"title"`
	Video    string `gorm:"type:varchar(50) not null" json:"video"`
	Message  string `gorm:"type:varchar(50) not null" json:"message"`
	Date     string `gorm:"type:varchar(50) not null" json:"date"`
	Image    string `gorm:"type:varchar(50) not null" json:"image"`
	Preacher string `gorm:"type:varchar(50) not null" json:"preacher"`
}
