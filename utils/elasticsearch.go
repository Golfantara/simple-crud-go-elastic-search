package utils

import (
	"crypto/tls"
	"elasticsearch/config"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticClient() (*elasticsearch.Client, error) {
	cfg := config.LoadDBConfig()
	log.Printf("Connecting to Elasticsearch: %s", cfg.ELASTIC_URL)
	log.Printf("Using user: %s", cfg.ELASTIC_USER)
	
	// Create a new Elasticsearch client
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.ELASTIC_URL},
		Username:  cfg.ELASTIC_USER,
		Password:  cfg.ELASTIC_PASS,
		Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
	})

	if err != nil {
		log.Printf("Error creating the ElasticSearch client: %s", err)
		return nil, err
	}

	log.Println("Successfully connected to Elasticsearch")

	return client, nil
}
