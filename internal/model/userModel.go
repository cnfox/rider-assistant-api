package model

// User 队员表
type User struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT"` // 号码牌
	Openid   int    `gorm:"type:int(11)"`               // 微信id
	Name     string `gorm:"type:varchar(255)"`          // 江湖称号
	Nickname string `gorm:"type:varchar(255)"`          // 昵称
	Avatar   string `gorm:"type:varchar(255)"`          // 头像
	Group    int    `gorm:"type:int(11)"`               // 分组
	Score    int    `gorm:"type:int(11)"`               // 积分
	IsLeader int    `gorm:"type:tinyint(4)"`            // 是否领队:1是、0否
}

// TableName 声明表名
func (User) TableName() string {
	return "user"
}
