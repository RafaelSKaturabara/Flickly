package viewmodel

import "github.com/google/uuid"

type ListAllSubcategoriesViewModel struct {
}

type SubcategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CategoryID  uuid.UUID `json:"categoryId"`
	Category    CategoryResponse `json:"category"`
}

type CategoryResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
}

type ListAllSubcategoriesViewModelResponse struct {
	Subcategories []SubcategoryResponse `json:"subcategories"`
}
