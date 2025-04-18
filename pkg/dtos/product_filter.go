package dtos

// ProductFilterDTO represents filter criteria for products
type ProductFilterDTO struct {
	SupplierId string   `form:"supplierId" json:"supplierId,omitempty"`
	CategoryId string   `form:"categoryId" json:"categoryId,omitempty"`
	MinPrice   *float64 `form:"minPrice" json:"minPrice,omitempty"`
	MaxPrice   *float64 `form:"maxPrice" json:"maxPrice,omitempty"`
	Rating     string   `form:"rating" json:"rating,omitempty"`
	SearchTerm string   `form:"searchTerm" json:"searchTerm,omitempty"`
	Tags       []string `form:"tags" json:"tags,omitempty"`
}
