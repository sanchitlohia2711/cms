package v1

//ProductCategoryMappingParams create update params
type ProductCategoryMappingParams struct {
	ProductId  uint        `json:"productId"`
	CategoryId uint        `json:"categoryId"`
	Tags       interface{} `json:"tags"`
}
