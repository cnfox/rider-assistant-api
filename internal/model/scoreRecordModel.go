package model

// ScoreRecord 积分流水表
type ScoreRecord struct {
	ID int `gorm:"primary_key;AUTO_INCREMENT"` // 主键
}

// TableName 声明表名
func (ScoreRecord) TableName() string {
	return "score_record"
}
