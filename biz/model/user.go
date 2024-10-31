package model

// GetActiveUserByID 用ID获取可登录用户
func GetActiveUserByID(ID interface{}) (User, error) {
	var user User
	//result := mysql.DB.Set("gorm:auto_preload", true).Where("status = ?", Active).First(&model_gen, ID)
	//return model_gen, result.Error
	return user, nil
}
