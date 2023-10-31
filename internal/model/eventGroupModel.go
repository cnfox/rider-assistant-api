package model

import "time"

// EventGroup 活动分组表
type EventGroup struct {
	ID         int64     `gorm:"primary_key;AUTO_INCREMENT;column:id" db:"id" json:"id" form:"id"`
	EventId    int64     `gorm:"column:event_id" db:"event_id" json:"event_id" form:"event_id"`             //  活动id
	Sequence   string    `gorm:"column:sequence" db:"sequence" json:"sequence" form:"sequence"`             //  分组序号
	LeaderId   int64     `gorm:"column:leader_id" db:"leader_id" json:"leader_id" form:"leader_id"`         //  领队id
	LeaderName string    `gorm:"column:leader_name" db:"leader_name" json:"leader_name" form:"leader_name"` //  领队花名
	Strength   string    `gorm:"column:strength" db:"strength" json:"strength" form:"strength"`             //  分组强度
	CreateAt   time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
	UpdateAt   time.Time `gorm:"column:update_at" db:"update_at" json:"update_at" form:"update_at"`
}

// TableName 声明表名
func (EventGroup) TableName() string {
	return "event_group"
}
