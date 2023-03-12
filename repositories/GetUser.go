package repositories

func GetUser(userid uint, user interface{}) error {
	return db.Where("id = ?", userid).First(user).Error
}
