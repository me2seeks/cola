package models

// UserAuth represents the user_auth table
type UserAuth struct {
	Base
	Version  int64  `gorm:"column:version;default:0;comment:版本号"`
	UserID   int64  `gorm:"column:user_id;default:0"`
	AuthKey  string `gorm:"column:auth_key;type:varchar(64);not null;default:'';comment:平台唯一id"`
	AuthType string `gorm:"column:auth_type;type:varchar(12);not null;default:'';comment:平台类型"`
}

func (UserAuth) TableName() string {
	return "user_auth"
}
