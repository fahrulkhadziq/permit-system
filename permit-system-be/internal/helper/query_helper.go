package helper

import "gorm.io/gorm"

func ApplyPagination(db *gorm.DB, page int, limit int) *gorm.DB {
	offset := (page - 1) * limit

	return db.Offset(offset).Limit(limit)
}

func ApplySorting(db *gorm.DB, sort string, order string) *gorm.DB {

	allowedSort := map[string]bool{
		"created_at":    true,
		"expired_at":    true,
		"document_name": true,
	}

	if !allowedSort[sort] {
		sort = "created_at"
	}

	if order != "asc" && order != "desc" {
		order = "desc"
	}

	return db.Order(sort + " " + order)
}

func ApplySearch(db *gorm.DB, search string, fields []string) *gorm.DB {

	if search == "" {
		return db
	}

	for i, field := range fields {
		query := field + " ILIKE ?"
		if i == 0 {
			db = db.Where(query, "%"+search+"%")
		} else {
			db = db.Or(query, "%"+search+"%")
		}
	}

	return db
}

func ApplyFilters(db *gorm.DB, filters map[string]interface{}) *gorm.DB {

	for key, value := range filters {
		if value == nil || value == "" {
			continue
		}
		db = db.Where(key+" = ?", value)
	}
	return db
}
