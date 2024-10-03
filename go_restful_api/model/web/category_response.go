package web

type CategoryResponse struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,min=3,max=200"`
}
