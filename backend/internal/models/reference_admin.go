package models

type CreateAuthorRequest struct {
	FullName string `json:"full_name"`
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}
