package models

type Image struct {
	Id         int    `json:"id" gorm:"-;primaryKey;AUTO_INCREMENT"`
	Info       string `json:"info"`
	Type       string `json:"type"`
	Source     string `json:"source"`
	Url        string `json:"url"`
	UpdateTime string `json:"update_time"`
	State      int    `json:"state"`
}

func (Image) TableName() string {
	return "images"
}
