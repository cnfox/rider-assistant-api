package model

// EventEnroll 活动报名签到表
type EventEnroll struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"` // 主键
}

// TableName 声明表名
func (EventEnroll) TableName() string {
	return "event_enroll"
}
