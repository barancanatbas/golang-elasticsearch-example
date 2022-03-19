package product

import (
	"context"
	"net/http"
	"shopping/internal/services/product"
	"shopping/requests"

	"github.com/gin-gonic/gin"
)

type IProduct interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Gets(c *gin.Context)
	Search(c *gin.Context)
}

type HandlerProduct struct {
	Servies product.ProductService
}

func NewHandler(services product.ProductService) HandlerProduct {
	return HandlerProduct{
		Servies: services,
	}
}

func (p HandlerProduct) Get(c *gin.Context) {
	ctx := context.Background()
	response, err := p.Servies.Get(&ctx)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}

	c.JSON(200, response)
}

func (p HandlerProduct) Gets(c *gin.Context) {
	ctx := context.Background()
	response, err := p.Servies.Gets(&ctx)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}

	c.JSON(200, response)
}

func (p HandlerProduct) Search(c *gin.Context) {
	param := c.Query("key")
	ctx := context.Background()
	response, err := p.Servies.Search(&ctx, param)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}

	c.JSON(200, response)
}

func (p HandlerProduct) Create(c *gin.Context) {
	ctx := context.Background()
	var req requests.Create

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := p.Servies.Create(req, &ctx)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, "success!")
}

func (p HandlerProduct) Delete(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	err := p.Servies.Delete(&ctx, id)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, "success!")
}

func (p HandlerProduct) Update(c *gin.Context) {
	ctx := context.Background()
	var req requests.Update

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := p.Servies.Update(req, &ctx)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, "success!")
}
