package models

type Users struct {
	Id int `gorm:"primaryKey" json:"id"`
	Roles string `json:"roles"`
	DescriptionRoles string `json:"description_roles"`
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	FullName string `json:"full_name"`
	Nis string `json:"nis"`
}
