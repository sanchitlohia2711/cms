package create

import (
	"context"
	"log"
)

//ProductMapping create the product mapping
func EntityMapping() (err error) {
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
					"meta":{
						"dynamic":"true",
						"type":"object",
						"enabled":"false"
					},
					"location":{
						"type":"geo_point"
					},
					"merchant_id":{
						"type":"text"
					},
					"city_id":{
						"type":"long"
					},
					"country_id":{
						"type":"long"
					},
					"type":{
						"type":"keyword"
					},
					"tags":{
						"type":"text",
						"fields":{
							"keyword":{
								"type":"keyword"
							}
						}
					},
					"active":{
						"type":"boolean"
					},
					"created_at":{
						"type":"date"
					},
					"updated_at":{
						"type":"date",
						"index":"false"
					}
				}
			}
		}
	}`

	// Use the IndexExists service to check if a specified index exists.
	ctx := context.Background()
	exists, err := esClient.IndexExists(conf.ES.EntityIndex).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	//if index does not exist, create a new one with the specified mapping
	if !exists {
		createIndex, err := esClient.CreateIndex(conf.ES.EntityIndex).BodyString(mapping).Do(ctx)
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
