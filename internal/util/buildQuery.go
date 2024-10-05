package util

import "gorm.io/gorm"

func BuildQuery(db *gorm.DB, equalConditions map[string]interface{}, likeConditions map[string]interface{}) *gorm.DB {
	// Add equal conditions
	for key, value := range equalConditions {
		db = db.Where(key, value)
	}

	// Add LIKE conditions
	for key, value := range likeConditions {
		db = db.Where(key+" LIKE ?", "%"+value.(string)+"%")
	}

	return db
}
