package models

type User struct {
	ID           string `json:"id" gorm:"primary_key"`
	NickName     string `json:"nickName" gorm:"unique"`
	HashPassword string
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	OtherName    string `json:"otherName"`
	Email        string `json:"email" gorm:"unique"`
	PhoneNumber  string `json:"phoneNumber" gorm:"unique"`
}
