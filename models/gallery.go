package models

type Gallery struct {
	BaseModel
	Title string `gorm:"type:varchar(50) not null" json:"title"`
	Image string `gorm:"type:varchar(50) not null" json:"image"`
	Date  string `gorm:"type:varchar(50) not null" json:"date"`
}
