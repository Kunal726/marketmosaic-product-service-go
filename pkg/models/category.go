package models

// Category represents the category entity in the database
type Category struct {
	CategoryId     uint       `gorm:"primaryKey;column:category_id"`
	CategoryName   string     `gorm:"column:category_name;not null"`
	ParentId       *uint      `gorm:"column:parent_id"`
	ParentCategory *Category  `gorm:"foreignKey:ParentId"`
	SubCategories  []Category `gorm:"foreignKey:ParentId"`
}

// TableName specifies the table name for the Category model
func (Category) TableName() string {
	return "marketmosaic_category"
}
