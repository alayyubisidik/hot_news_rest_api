package web

type CategoryCreateRequest struct {
	Name string `validate:"required,min=3,max=50" json:"name"`
}
