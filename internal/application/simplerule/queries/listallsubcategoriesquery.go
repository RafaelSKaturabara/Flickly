package queries

import (
	"context"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/core"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/entities"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/google/uuid"
)

type ListAllSubcategoriesQuery struct {
	UserSimpleRuleID uuid.UUID
}

type ListAllSubcategoriesQueryHandler struct {
	mediator              mediator.Mediator
	subcategoryRepository repositories.ISubcategoryRepository
}

func NewListAllSubcategoriesQueryHandler(serviceCollection utilities.IServiceCollection) *ListAllSubcategoriesQueryHandler {
	return &ListAllSubcategoriesQueryHandler{
		mediator:              utilities.GetService[mediator.Mediator](serviceCollection),
		subcategoryRepository: utilities.GetService[repositories.ISubcategoryRepository](serviceCollection),
	}
}

func (h *ListAllSubcategoriesQueryHandler) Handle(c context.Context, request mediator.Request) (mediator.Response, error) {
	query := request.(ListAllSubcategoriesQuery)

	subcategories, err := h.subcategoryRepository.Find(func(s *entities.Subcategory) bool {
		return s.UserSimpleRuleID == nil || *s.UserSimpleRuleID == query.UserSimpleRuleID
	})

	if err != nil {
		return nil, core.DoNotUseThisGenericError(err)
	}

	return subcategories, nil
}
