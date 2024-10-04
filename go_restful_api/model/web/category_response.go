package web

type CategoryResponse struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,min=3,max=200" json:"name"`
}
