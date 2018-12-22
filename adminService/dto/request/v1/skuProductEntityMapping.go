package v1

//SkuProductEntityMappingParams create update sku product entity mapping params
type SkuProductEntityMappingParams struct {
	ProductID uint `json:"productId"`
	EntityID  uint `json:"entityId"`
	SkuID     uint `json:"skuId"`
}
