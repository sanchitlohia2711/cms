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
	vertical, err := createVertical()
	if err != nil {
		log.Fatalf(err.Error())
	}
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
			_, err := createSalt(vertical, scanner.Text())
			if err != nil {
				log.Fatal(err.Error())
			}
			continue
		}
		_, err := createComposition(vertical, scanner.Text())
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func createSalt(vertical *adminModel.Vertical, salt string) (*adminModel.Category, error) {
	categoryParams := &requestDTOV1.CategoryParams{}
	categoryParams.Name = salt
	categoryParams.Description = "Desciption of" + salt
	categoryParams.VerticalID = vertical.ID
	return adminModel.CreateCategory(categoryParams)
}

func createComposition(vertical *adminModel.Vertical, composition string) (*adminModel.Product, error) {
	split := strings.Split(composition, ",")
	brand, err := createBrand(split[3])
	if err != nil {
		log.Fatalf("Error occured" + err.Error())
	}

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
	productParams.Attributes = make(map[string]interface{})
	productParams.Attributes["quantity"] = split[4]
	return adminModel.CreateProduct(productParams)
}

func createVertical() (*adminModel.Vertical, error) {
	verticalParams := &requestDTOV1.VerticalParams{}
	verticalParams.Name = "Medical"
	verticalParams.Description = "Used for medical purpose"
	vertical, err := adminModel.CreateVertical(verticalParams)
	if err != nil {
		return nil, err
	}
	return vertical, nil
}

func createBrand(name string) (brand *adminModel.Brand, err error) {
	brandParams := &requestDTOV1.BrandParams{}
	brandParams.Name = name
	brandParams.Description = "brandDescription"
	return adminModel.CreateBrandIfNotExists(brandParams)
}
