package model

import "time"

// ScoreRecord 积分流水表
type ScoreRecord struct {
	ID       int64     `gorm:"primary_key;AUTO_INCREMENT;column:id" db:"id" json:"id" form:"id"`
	Uid      int64     `gorm:"column:uid" db:"uid" json:"uid" form:"uid"`                     //  队员id
	Type     int64     `gorm:"column:type" db:"type" json:"type" form:"type"`                 //  加积分原因:1活动签到、2历史积分、3队服、4全程、5文章、6视频
	Score    int64     `gorm:"column:score" db:"score" json:"score" form:"score"`             //  积分
	Note     string    `gorm:"column:note" db:"note" json:"note" form:"note"`                 //  备注
	EventId  int64     `gorm:"column:event_id" db:"event_id" json:"event_id" form:"event_id"` //  活动id
	AdminId  int64     `gorm:"column:admin_id" db:"admin_id" json:"admin_id" form:"admin_id"` //  管理员id
	CreateAt time.Time `gorm:"column:create_at" db:"create_at" json:"create_at" form:"create_at"`
}

// TableName 声明表名
func (ScoreRecord) TableName() string {
	return "score_record"
}
