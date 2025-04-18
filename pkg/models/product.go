package models

import (
	"time"
)

// Product represents the product entity in the database
type Product struct {
	ProductId     uint      `gorm:"primaryKey;column:product_id;autoIncrement"`
	ProductName   string    `gorm:"column:product_name;not null"`
	Description   string    `gorm:"column:description"`
	Price         float64   `gorm:"column:price;not null"`
	StockQuantity int       `gorm:"column:stock_quantity"`
	SupplierId    uint      `gorm:"column:supplier_id;not null"`
	Supplier      User      `gorm:"foreignKey:SupplierId;references:Id"`
	ImageUrl      string    `gorm:"column:image_url"`
	DateAdded     time.Time `gorm:"column:date_added;not null"`
	IsActive      bool      `gorm:"column:is_active;not null"`
	Rating        string    `gorm:"column:rating"`
	CategoryId    uint      `gorm:"column:category_id;not null"`
	Category      Category  `gorm:"foreignKey:CategoryId;references:CategoryId"`
	Tags          []Tag     `gorm:"many2many:product_tags;foreignKey:ProductId;joinForeignKey:product_id;References:TagId;joinReferences:tag_id"`
}

// TableName specifies the table name for the Product model
func (Product) TableName() string {
	return "marketmosaic_product"
}
