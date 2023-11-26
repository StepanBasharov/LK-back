package schemas

type SigUpUser struct {
	NickName    string `json:"nickName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	OtherName   string `json:"otherName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Password2   string `json:"password2"`
}

type LoginUser struct {
	NickName string `json:"nickName"`
	Password string `json:"password"`
}
