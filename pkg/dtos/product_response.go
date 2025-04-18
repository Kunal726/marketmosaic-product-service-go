package dtos

import "github.com/Kunal726/market-mosaic-common-lib-go/pkg/dtos"

// ProductResponseDTO represents the response for product-related operations
type ProductResponseDTO struct {
	*dtos.BaseResponseDTO
	Product     *ProductDetailsDTO   `json:"product,omitempty"`
	ProductList []*ProductDetailsDTO `json:"productList,omitempty"`
}
