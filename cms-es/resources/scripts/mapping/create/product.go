package create

import (
	"context"
	"log"

	"github.com/ko/cms-db/cms-es/es"
)

var (
	esClient = es.GetClient()
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
					"categories":{
						"type":"long"
					},
					"tags" : {
						"type":"text"
					},
					"tnc":{
						"type":"text",
						"index":"false"
					}
				}
			}
		}
	}`

	// Use the IndexExists service to check if a specified index exists.
	ctx := context.Background()
	exists, err := esClient.IndexExists("product").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	//if index does not exist, create a new one with the specified mapping
	if !exists {
		createIndex, err := esClient.CreateIndex("product").BodyString(mapping).Do(ctx)
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
