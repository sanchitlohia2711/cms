package es

import "github.com/olivere/elastic"

var (
	client = &elastic.Client{}
)

func init() {
	var err error
	client, err = elastic.NewClient()
	if err != nil {
		// Handle error
	}
}

//GetClient returns the elastic search client
func GetClient() *elastic.Client {
	return client
}
