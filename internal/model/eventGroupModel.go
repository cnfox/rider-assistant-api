package model

// EventGroup 活动分组表
type EventGroup struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"` // 主键
}

// TableName 声明表名
func (EventGroup) TableName() string {
	return "event_group"
}
