package request

// nambahin validate pake validator v10
type RequestCreateProduct struct {
	Name string `json:"name" validate:"required,min=1,max=200"`
}

type RequestUpdateProduct struct {
	Id   int    `json:"id" validate:"required,numeric"`
	Name string `json:"name" validate:"required,min=1,max=200"`
}

type RequestDeleteProduct struct {
	Id int `json:"id" validate:"required,numeric"`
}
