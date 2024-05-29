package models

type Role struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}

func (Role) TableName() string {
	return "roles"
}
