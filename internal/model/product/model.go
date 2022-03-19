package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"shopping/internal/elasticsearch"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type IProduct interface {
	Get(products *[]Product, ctx *context.Context) error
	Create(product []byte, id string, ctx *context.Context) error
	Update(product string, id string, ctx *context.Context) error
	Delete(id string, ctx *context.Context) error
	Gets(ctx *context.Context) ([]*Product, error)
	Search(datajson []byte, ctx *context.Context) ([]*Product, error)
}

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"Price"`
	Stock int    `json:"stock"`
	Color string `json:"color"`
}

type ModelProduct struct {
	Elastic elasticsearch.ElasticSearch
}

func NewModel(elastic elasticsearch.ElasticSearch) ModelProduct {
	return ModelProduct{
		Elastic: elastic,
	}
}

type document struct {
	Source interface{} `json:"_source"`
}

type productResponse struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []struct {
			Source *Product `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (p ModelProduct) Get(products *Product, ctx *context.Context) error {

	req := esapi.GetRequest{
		Index:      "product",
		DocumentID: "1",
	}

	res, err := req.Do(*ctx, p.Elastic.Client)
	defer res.Body.Close()

	var body document

	body.Source = products
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return err
	}

	return err
}

func (p ModelProduct) Gets(ctx *context.Context) ([]*Product, error) {
	req := esapi.SearchRequest{
		Index: []string{"product"},
	}

	res, err := req.Do(*ctx, p.Elastic.Client)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body productResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}

	products := make([]*Product, len(body.Hits.Hits))
	for i, v := range body.Hits.Hits {
		products[i] = v.Source
	}

	return products, err
}

func (p ModelProduct) Search(datajson []byte, ctx *context.Context) ([]*Product, error) {

	// var buf bytes.Buffer
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match": map[string]interface{}{
	// 			"color": "red",
	// 		},
	// 	},
	// }
	// if err := json.NewEncoder(&buf).Encode(query); err != nil {
	// 	log.Fatalf("Error encoding query: %s", err)
	// }
	// req := esapi.SearchRequest{
	// 	Index: []string{"product"},
	// }

	req := esapi.SearchRequest{
		Index: []string{"product"},
		Body:  bytes.NewReader(datajson),
	}

	res, err := req.Do(*ctx, p.Elastic.Client)

	var body productResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, err
	}

	products := make([]*Product, len(body.Hits.Hits))
	for i, v := range body.Hits.Hits {
		products[i] = v.Source
	}

	return products, err

	// // Perform the search request.
	// res, err := p.Elastic.Client.Search(
	// 	p.Elastic.Client.Search.WithContext(*ctx),
	// 	p.Elastic.Client.Search.WithIndex("product"),
	// 	p.Elastic.Client.Search.WithBody(&buf),
	// 	p.Elastic.Client.Search.WithTrackTotalHits(true),
	// 	p.Elastic.Client.Search.WithPretty(),
	// )
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }
	// defer res.Body.Close()
}

func (p ModelProduct) Create(product []byte, id string, ctx *context.Context) error {

	req := esapi.CreateRequest{
		Index:      "product",
		DocumentID: id,
		Body:       bytes.NewReader(product),
	}

	res, err := req.Do(*ctx, p.Elastic.Client)
	defer res.Body.Close()

	return err
}

func (p ModelProduct) Update(product []byte, id string, ctx *context.Context) error {
	fmt.Println(id)
	req := esapi.UpdateRequest{
		Index:      "product",
		DocumentID: id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, string(product)))),
	}

	res, err := req.Do(*ctx, p.Elastic.Client)
	defer res.Body.Close()

	return err
}

func (p ModelProduct) Delete(id string, ctx *context.Context) error {

	req := esapi.DeleteRequest{
		Index:      "product",
		DocumentID: id,
	}

	res, err := req.Do(*ctx, p.Elastic.Client)
	defer res.Body.Close()

	return err
}
