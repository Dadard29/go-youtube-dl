package models

import "time"

type Status struct {
	StatusId int `gorm:"type:int;index:status_id;primary_key;auto_increment"`
	Token string `gorm:"type:varchar(70);index:token;"`
	DateStarted time.Time `gorm:"type:datetime;index:date_started"`
	DateFinished time.Time `gorm:"type:datetime;index:date_finished"`
	Progress int `gorm:"type:int;index:progress"`
	Done bool `gorm:"type:bool;index:done"`
	Message string `gorm:"type:varchar(30);index:message"`
}

func (Status) TableName() string {
	return "status"
}

type StatusJson struct {
	DateStarted time.Time
	DateFinished time.Time
	Progress int
	Done bool
	Message string
}

func NewStatusJson(s Status) StatusJson {
	return StatusJson{
		DateStarted: s.DateStarted,
		DateFinished: s.DateFinished,
		Progress: s.Progress,
		Done:     s.Done,
		Message:  s.Message,
	}

}