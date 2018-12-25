package create

import (
	"context"
	"log"
)

//SkuMapping create the product mapping
func SkuMapping() (err error) {
	var mapping = `
	{
		"settings":{
			"number_of_shards": 1,
			"number_of_replicas": 0
		},
		"mappings":{
			"default":{
				"dynamic":"strict",
				"properties":{
					"name":{
						"type":"text"
					},
					"description":{
						"type":"text"
					},
					"attributes":{
						"dynamic":"true",
						"type":"object",
						"enabled":"false"
					},
					"tnc":{
						"type":"text",
						"index":"false"
					},
					"how_to_redeem":{
						"type":"text",
						"index":"false"
					},
					"terms_conditions":{
						"type":"text",
						"index":"false"
					},
					"return_policy":{
						"type":"text",
						"index":"false"
					},
					"vertical_id":{
						"type":"long"
					},
					"brand_id":{
						"type":"long"
					},
					"price":{
						"type":"long"
					},
					"offer_price":{
						"type":"long"
					}
					"status":{
						"type":"integer"
					},
					"visibility":{
						"type":"integer"
					},
					"start_date":{
						"type":"date"
					},
					"end_date":{
						"type":"date"
					},
					"created_at":{
						"type":"date"
					},
					"updated_at":{
						"type":"date",
						"index":"false"
					},
					"share_url":{
						"type":"keyword",
						"index":"false"
					},
					"image_url":{
						"type":"keyword",
						"index":"false"
					},
					"thumb_url":{
						"type":"keyword",
						"index":"false"
					},
					"merchant_id":{
						"type":"long"
					},
					"input_fields":{
						"dynamic":"true",
						"type":"object",
						"enabled":"false"
					},
					"pay_type_supported":{
						"type":"object",
						"dynamic":"true"
						"enabled":"false"
					}
				}
			}
		}
	}`

	// Use the IndexExists service to check if a specified index exists.
	ctx := context.Background()
	exists, err := esClient.IndexExists(conf.ES.SkuIndex).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	//if index does not exist, create a new one with the specified mapping
	if !exists {
		createIndex, err := esClient.CreateIndex(conf.ES.SkuIndex).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			log.Println(createIndex)
		} else {
			log.Println("successfully created index")
		}
	} else {
		log.Println("Index already exist")
	}
	return
}
