package main

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var ES *elasticsearch.TypedClient

func ConnectES() {
	var err error
	ES, err = elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	_, err = ES.Info().Do(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to Elasticsearch: %s", err)
	}

	log.Println("Successfully connected to Elasticsearch")
}
