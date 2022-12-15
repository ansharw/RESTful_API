package response

type ResponseProduct struct {
	Id int `json:"id"`
	Name string `json:"name"`
	CategoryId int `json:"category_id"`
}