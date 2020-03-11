package models

type VideoModel struct {
	VideoId string `gorm:"type:varchar(30);index:video_id;primary_key"`
	Token string `gorm:"type:varchar(30);index:token"`
	Title string `gorm:"type:varchar(30);index:title"`
	Album string `gorm:"type:varchar(30);index:album"`
	Artist string `gorm:"type:varchar(30);index:artist"`
	Date string `gorm:"type:date;index:date"`
}

func (VideoModel) TableName() string {
	return "video"
}
