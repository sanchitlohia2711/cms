package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	adminModel "github.com/ko/cms-db/adminService/adminModel"
	"github.com/ko/cms-db/adminService/dto"
	requestDTOV1 "github.com/ko/cms-db/adminService/dto/request/v1"
)

func main() {
	vertical := createVertical()
	file, err := os.Open("./genmed.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		if strings.Contains(scanner.Text(), "Brand Name") {
			continue
		}
		if len(strings.Split(scanner.Text(), ",")) == 1 {
			createSalt(vertical, scanner.Text())
			continue
		}
		createComposition(vertical, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createSalt(vertical *adminModel.Vertical, salt string) {

	fmt.Println(salt)
}

func createComposition(vertical, composition string) {
	split = strings.split(composition, ",")
	brand := createBrand(brandName)

	meta := dto.Meta{}
	meta.Title = "someTitle"
	meta.Description = "someDescription"
	productParams := &requestDTOV1.ProductParams{}
	productParams.Name = split[0]
	productParams.Description = split[1]
	productParams.VerticalID = vertical.ID
	productParams.BrandID = brand.ID
	productParams.StartDate = time.Now()
	productParams.EndDate = time.Now()
	productParams.ImageURL = "http://image"
	productParams.ShareURL = "http://share"
	productParams.ThumbURL = "http://thumb"
	productParams.InputFields = "sdasdsa"
	productParams.HowToRedeem = meta
	productParams.TermsConditions = meta
	productParams.ReturnPolicy = meta
	productParams.Tags = []string{"sun", "moon"}
	return CreateProduct(productParams)
	fmt.Println(composition)
}

func createCategory(vertical *adminModel.Vertical, categoryName string) (category *adminModel.Category, err error) {
	categoryParams := &requestDTOV1.CategoryParams{}
	categoryParams.Name = categoryName
	categoryParams.Description = "Desciption of" + categoryName
	categoryParams.VerticalID = vertical.ID
	return adminModel.CreateCategory(categoryParams)
}

func createProduct(vertical *adminModel.CreateCountryVertical, brand *adminModel.Brand) (product *adminModel.Product, err error) {

}

func createVertical() {
	verticalParams := &requestDTOV1.VerticalParams{}
	verticalParams.Name = "Medical"
	verticalParams.Description = "Used for medical purpose"
	adminModel.CreateVertical(verticalParams)
	return
}

func createBrand(name string) (brand *Brand, err error) {
	brandParams := &requestDTOV1.BrandParams{}
	brandParams.Name = name
	brandParams.Description = "brandDescription"
	return adminModel.CreateBrandIfNotExists(brandParams)
}
