package models

type NewsletterSubscription struct {
	BaseModel
	Email string `gorm:"type:varchar(50) not null" json:"email"`
}
