package models

type User struct {
	ID           string `json:"id" gorm:"primary_key"`
	NickName     string `json:"nickName"`
	HashPassword string
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	OtherName    string `json:"otherName"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
}
