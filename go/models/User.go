package models

type User struct {
	Id    int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name  string `json:"name"`
	Token string `json:"token"`
	Role  int    `json:"role"`
	State int    `json:"state"`
}

func (User) TableName() string {
	return "users"
}
