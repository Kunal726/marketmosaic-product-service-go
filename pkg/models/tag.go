package models

// Tag represents the tag entity in the database
type Tag struct {
	TagId    uint      `gorm:"primaryKey;column:tag_id"`
	TagName  string    `gorm:"column:tag_name;not null"`
	Products []Product `gorm:"many2many:product_tags"`
}

// TableName specifies the table name for the Tag model
func (Tag) TableName() string {
	return "marketmosaic_tag"
}
