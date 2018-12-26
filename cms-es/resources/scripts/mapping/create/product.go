package create

import (
	"context"
	"log"
)

//ProductMapping create the product mapping
func ProductMapping() (err error) {
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
					"category":{
						"type":"long"
					},
					"tags" : {
						"type" : "text",
						"fields":{
							"keyword":{
								"type":"keyword"
							}
						}
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
	exists, err := esClient.IndexExists(conf.ES.ProductIndex).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	//if index does not exist, create a new one with the specified mapping
	if !exists {
		createIndex, err := esClient.CreateIndex(conf.ES.ProductIndex).BodyString(mapping).Do(ctx)
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
