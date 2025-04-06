package models

type Blog struct {
	Id        uint   `json:"id"`
	Title	  string `json:"title"`
	Description	string `json:"description"`
	Image string `json:"image"`
	UserId uint `json:"userId"`
	User User `json:"user"; gorm:"foreignkey:UserId"`
}
