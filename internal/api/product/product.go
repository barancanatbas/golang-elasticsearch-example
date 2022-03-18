package product

import (
	"shopping/internal/elasticsearch"
	handler "shopping/internal/handler/product"
	model "shopping/internal/model/product"
	service "shopping/internal/services/product"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, e *elasticsearch.ElasticSearch) {

	db := model.NewModel(*e)
	service := service.NewService(db)
	handler := handler.NewHandler(service)

	r.GET("/product", handler.Get)
	r.GET("/products", handler.Gets)
	r.POST("/product", handler.Create)
	r.PUT("/product", handler.Update)
	r.DELETE("/product/:id", handler.Delete)
}
