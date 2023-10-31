package model

import "time"

// Event 活动表
type Event struct {
	ID           int64     `gorm:"primary_key;AUTO_INCREMENT;column:id" db:"id" json:"id" form:"id"`
	Title        string    `gorm:"column:title" db:"title" json:"title" form:"title"`                                 //  活动标题
	LeaderId     int64     `gorm:"column:leader_id" db:"leader_id" json:"leader_id" form:"leader_id"`                 //  领队id
	LeaderName   string    `gorm:"column:leader_name" db:"leader_name" json:"leader_name" form:"leader_name"`         //  领队花名
	EventDate    string    `gorm:"column:event_date" db:"event_date" json:"event_date" form:"event_date"`             //  活动日期
	StartTime    string    `gorm:"column:start_time" db:"start_time" json:"start_time" form:"start_time"`             //  集合时间
	EventRoute   string    `gorm:"column:event_route" db:"event_route" json:"event_route" form:"event_route"`         //  活动路线
	Strength     string    `gorm:"column:strength" db:"strength" json:"strength" form:"strength"`                     //  活动强度
	EstimateTime string    `gorm:"column:estimate_time" db:"estimate_time" json:"estimate_time" form:"estimate_time"` //  预计耗时
	Notice       string    `gorm:"column:notice" db:"notice" json:"notice" form:"notice"`                             //  特别说明
	Code         string    `gorm:"column:code" db:"code" json:"code" form:"code"`                                     //  签到口令
	Score        int64     `gorm:"column:score" db:"score" json:"score" form:"score"`                                 //  签到积分
	UpperLimit   int64     `gorm:"column:upper_limit" db:"upper_limit" json:"upper_limit" form:"upper_limit"`         //  人数上限
	CreateAt     time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
	UpdateAt     time.Time `gorm:"column:update_at" db:"update_at" json:"update_at" form:"update_at"`
}

// TableName 声明表名
func (Event) TableName() string {
	return "event"
}
