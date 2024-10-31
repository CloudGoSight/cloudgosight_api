package model

import "gorm.io/gorm"

// User 用户模型
type User struct {
	// 表字段
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index"`
	Nick      string `gorm:"size:50"`
	Password  string `json:"-"`
	Status    int
	GroupID   uint
	Storage   uint64
	TwoFactor string
	Avatar    string
	Options   string `json:"-" gorm:"size:4294967295"`
	Authn     string `gorm:"size:4294967295"`

	// 关联模型
	// 用户组
	//Group  Group  `gorm:"save_associations:false:false"`
	// 存储
	//Policy Policy `gorm:"PRELOAD:false,association_autoupdate:false"`

	// 数据库忽略字段
	// 个性化
	//OptionsSerialized UserOption `gorm:"-"`
}

// GetActiveUserByID 用ID获取可登录用户
func GetActiveUserByID(ID interface{}) (User, error) {
	var user User
	//result := mysql.DB.Set("gorm:auto_preload", true).Where("status = ?", Active).First(&model_gen, ID)
	//return model_gen, result.Error
	return user, nil
}
