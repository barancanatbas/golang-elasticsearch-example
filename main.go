package main

import (
	"log"
	"shopping/internal/api/product"
	"shopping/internal/elasticsearch"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Bootstrap elasticsearch.
	elastic, err := elasticsearch.New([]string{"http://0.0.0.0:9200"})
	if err != nil {
		log.Fatalln(err)
	}

	product.Init(r, elastic)

	// Start HTTP server.
	r.Run()
}
