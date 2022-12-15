package domain

import "belajar-rest-api/model/response"

type Product struct {
	id         int
	categoryId int
	name       string
}

func (p *Product) ToResponseProduct() response.ResponseProduct {
	return response.ResponseProduct{
		Id:         p.id,
		Name:       p.name,
		CategoryId: p.categoryId,
	}
}

func (p *Product) SetId(id *int) {
	p.id = *id
}

func (p *Product) SetCategoryId(categoryId *int) {
	p.categoryId = *categoryId
}

func (p *Product) SetName(name *string) {
	p.name = *name
}

func (p *Product) GetId() *int {
	return &p.id
}

func (p *Product) GetCategryId() *int {
	return &p.categoryId
}

func (p *Product) GetName() *string {
	return &p.name
}
