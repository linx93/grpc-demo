package service

import (
	"context"
	"log"
)

var ProductService = &productService{}

type productService struct {
}

func (p *productService) GetProductStock(context context.Context, request *Request) (*Response, error) {
	stock := p.GetStockById(request.Id)
	return &Response{Stock: stock}, nil
}

func (p *productService) GetStockById(id uint32) uint32 {
	//通过id去查询库存
	stock := id + 1

	log.Printf("查询了一次库存【stock=%v】", stock)
	return stock
}
