package elasticsearch

import "github.com/elastic/go-elasticsearch/v7"

// type ElasticSearch struct {
// 	Client *elasticsearch.Client
// }

// func New() (ElasticSearch, error) {

// 	client, err := elasticSearch.NewClient(elastic.SetURL("http://localhost:9200"),
// 		elastic.SetSniff(false),
// 		elastic.SetHealthcheck(false))

// 	fmt.Println("ES initialized...")

// 	return client, err
// }

type ElasticSearch struct {
	Client *elasticsearch.Client
}

func New(addresses []string) (*ElasticSearch, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticSearch{
		Client: client,
	}, nil
}
