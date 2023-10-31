package model

import "time"

// EventEnroll  活动报名签到表
type EventEnroll struct {
	ID       int64     `gorm:"primary_key;AUTO_INCREMENT;column:id" db:"id" json:"id" form:"id"`
	Uid      int64     `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`                     //  用户id
	EventId  int64     `gorm:"column:event_id" db:"event_id" json:"event_id" form:"event_id"` //  活动id
	CheckIn  int64     `gorm:"column:check_in" db:"check_in" json:"check_in" form:"check_in"` //  是否签到1是0否
	CreateAt time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at" db:"update_at" json:"update_at" form:"update_at"`
}

// TableName 声明表名
func (EventEnroll) TableName() string {
	return "event_enroll"
}
