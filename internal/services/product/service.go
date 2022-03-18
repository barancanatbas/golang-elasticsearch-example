package product

import (
	"context"
	"encoding/json"
	"shopping/internal/model/product"
	"shopping/requests"

	"github.com/google/uuid"
)

type IProduct interface {
	Get(ctx *context.Context) ([]product.Product, error)
	Create(req requests.Create, ctx *context.Context) error
	Update(req requests.Update, ctx *context.Context) error
	Delete()
	Gets(ctx *context.Context) ([]*product.Product, error)
}

type ProductService struct {
	Db product.ModelProduct
}

func NewService(db product.ModelProduct) ProductService {
	return ProductService{
		Db: db,
	}
}

func (p ProductService) Get(ctx *context.Context) (product.Product, error) {
	var products product.Product

	err := p.Db.Get(&products, ctx)

	return products, err
}

func (p ProductService) Gets(ctx *context.Context) ([]*product.Product, error) {
	data, err := p.Db.Gets(ctx)

	return data, err
}

func (p ProductService) Create(req requests.Create, ctx *context.Context) error {
	id := uuid.NewString()
	product := product.Product{
		Name:  req.Name,
		Color: req.Color,
		Price: req.Price,
		Stock: req.Stock,
		ID:    id,
	}

	datajson, _ := json.Marshal(&product)

	err := p.Db.Create(datajson, id, ctx)

	return err
}

func (p ProductService) Update(req requests.Update, ctx *context.Context) error {
	product := product.Product{
		Name:  req.Name,
		Color: req.Color,
		Price: req.Price,
		Stock: req.Stock,
	}

	datajson, _ := json.Marshal(&product)

	err := p.Db.Update(datajson, req.ID, ctx)

	return err
}

func (p ProductService) Delete() {

}
