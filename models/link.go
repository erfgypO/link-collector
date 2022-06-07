package models

type Link struct {
	Id      uint   `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Url     string `json:"url"`
	UserID  uint   `json:"-"`
	Created int64  `json:"created" gorm:"autoCreateTime"`
	Updated int64  `json:"updated" gorm:"autoUpdateTime"`
}

type LinkDto struct {
	Name string `json:"name" binding:"required"`
	Url  string `json:"url" binding:"required"`
}

func CreateLink(link *Link) {
	DB.Create(link)
}

func GetLinks(userId uint) []Link {
	var links []Link

	DB.Where("user_id = ?", userId).Find(&links)

	return links
}

func DeleteLink(id uint, userId uint) {
	link := Link{
		Id: id,
	}
	DB.Where("user_id = ?", userId).Delete(&link)
}
