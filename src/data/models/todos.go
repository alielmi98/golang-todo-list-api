package models

type Todo struct {
	BaseModel
	Title       string `gorm:"size=70;type:string;not null;"`
	Description string `gorm:"size=200;type:string;not null;"`
	Completed   bool   `gorm:"default:false"`
	User        User   `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId      int    `gorm:"not null"`
}
