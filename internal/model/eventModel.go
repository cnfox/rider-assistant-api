package model

// Event 活动表
type Event struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"` // 主键
}

// TableName 声明表名
func (Event) TableName() string {
	return "event"
}
