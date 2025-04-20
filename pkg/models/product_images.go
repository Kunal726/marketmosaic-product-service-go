package models

type ProductImage struct {
	ImageID   uint   `gorm:"primaryKey"`
	ProductID uint   `gorm:"column:product_id;not null;index"`
	URL       string `gorm:"column:url;not null"`
	AltText   string `gorm:"column:alt_text"`
	IsPrimary bool   `gorm:"column:is_primary"`
}

func (ProductImage) TableName() string {
	return "marketmosaic_product_images"
}
