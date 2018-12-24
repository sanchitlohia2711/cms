package adminmodel

import (
	"fmt"
	"time"

	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
	"github.com/ko/cms-db/model"
)

func TestModel() (err error) {
	country, err := TestCountry()
	if err != nil {
		return
	}
	city, err := TestCity(country)
	if err != nil {
		return
	}

	brand, err := TestBrand()
	if err != nil {
		return
	}

	vertical, err := TestVertical()
	if err != nil {
		return
	}

	if err = CreateBrandVericalMapping(brand.ID, vertical.ID); err != nil {
		return
	}

	category, err := TestCategory(vertical)
	if err != nil {
		return
	}

	product, err := TestProduct(vertical, brand)
	if err != nil {
		return
	}

	categories := []*Category{category}
	err = product.AssociateCategories(categories)
	if err != nil {
		return
	}

	entity, err := TestEntity(city, country)
	if err != nil {
		return
	}

	err = TestPerson(entity)
	if err != nil {
		return
	}

	event, err := TestEvent(entity)
	if err != nil {
		return
	}

	sku, err := TestSku(product, event)
	if err != nil {
		return
	}

	err = TestSkuProductEntityMapping(product, entity, sku)
	if err != nil {
		return
	}

	skuProductEntityMapping := &model.SkuProductEntityMapping{}
	skuProductEntityMapping.ProductID = product.ID
	s, err := skuProductEntityMapping.Get()
	if err != nil {
		return
	}
	fmt.Println(s)
	return
	// TestBrandVerticalMapping()
	// TestCategory()
	// TestCategoryFilters()

}

func TestCountry() (country *Country, err error) {
	countryParams := &requestDTOV1.CountryParams{}
	countryParams.Name = "country"
	countryParams.Description = "countryDescription"
	return CreateCountry(countryParams)
}

func TestCity(country *Country) (city *City, err error) {
	cityParams := &requestDTOV1.CityParams{}
	cityParams.Name = "city"
	cityParams.Description = "cityDescription"
	cityParams.CountryID = country.ID
	return CreateCity(cityParams)
}

func TestBrand() (brand *Brand, err error) {
	brandParams := &requestDTOV1.BrandParams{}
	brandParams.Name = "brandName"
	brandParams.Description = "brandDescription"
	return CreateBrand(brandParams)
}

func TestVertical() (vertical *Vertical, err error) {
	verticalParams := &requestDTOV1.VerticalParams{}
	verticalParams.Name = "verticalName"
	verticalParams.Description = "verticalDescription"
	return CreateVertical(verticalParams)
}

func TestCategory(vertical *Vertical) (category *Category, err error) {
	categoryParams := &requestDTOV1.CategoryParams{}
	categoryParams.Name = "category"
	categoryParams.Description = "description"
	categoryParams.VerticalID = vertical.ID
	return CreateCategory(categoryParams)
}

func TestProduct(vertical *Vertical, brand *Brand) (product *Product, err error) {
	productParams := &requestDTOV1.ProductParams{}
	productParams.Name = "product"
	productParams.Description = "productDescription"
	productParams.VerticalID = vertical.ID
	productParams.BrandID = brand.ID
	return CreateProduct(productParams)
}

func TestEntity(city *City, country *Country) (entity *Entity, err error) {
	entityParams := &requestDTOV1.EntityParams{}
	entityParams.CityID = city.ID
	entityParams.CountryID = country.ID
	entityParams.Description = "entityDescription"
	entityParams.Type = "sfdsf"
	entityParams.Lat = 4.23
	entityParams.Lon = 2.54
	entityParams.Name = "entityName"
	entityParams.MerchantID = 1233
	return CreateEntity(entityParams)
}

func TestPerson(entity *Entity) (err error) {
	personParams := &requestDTOV1.PersonParams{}
	personParams.Name = "person"
	personParams.Description = "personDescription"
	return CreatePerson(personParams)
}

func TestEvent(entity *Entity) (event *Event, err error) {
	eventParams := &requestDTOV1.EventParams{}
	eventParams.Name = "event"
	eventParams.Description = "eventDescription"
	eventParams.StartTime = time.Now()
	eventParams.EndTime = time.Now()
	eventParams.StartDate = time.Now()
	eventParams.EntityID = entity.ID
	return CreateEvent(eventParams)
}

func TestSku(product *Product, event *Event) (sku *Sku, err error) {
	skuParams := &requestDTOV1.SkuParams{}
	skuParams.Price = 1
	skuParams.OfferPrice = 2
	skuParams.Name = "sku"
	skuParams.Description = "skuDescription"
	skuParams.ProductID = product.ID
	skuParams.EventID = event.ID
	skuParams.LongRichDesc = "longrich"
	skuParams.PayTypeSupported = "CC"
	skuParams.MerchantID = "123"
	skuParams.VerticalID = product.VerticalID
	return CreateSku(skuParams)
}

func TestSkuProductEntityMapping(product *Product, entity *Entity, sku *Sku) (err error) {
	skuProductEntityMappingParams := &requestDTOV1.SkuProductEntityMappingParams{}
	skuProductEntityMappingParams.ProductID = product.ID
	skuProductEntityMappingParams.EntityID = entity.ID
	skuProductEntityMappingParams.SkuID = sku.ID
	return sku.AssociateEntity(skuProductEntityMappingParams)
}
