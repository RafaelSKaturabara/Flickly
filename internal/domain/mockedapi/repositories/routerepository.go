package repositories	

import (
	"context"

	"github.com/RafaelSKaturabara/Flickly/internal/domain/mockedapi/entities"
	"github.com/google/uuid"
)

type IRouteRepository interface {
	// Métodos básicos
	CreateRoute(ctx context.Context, route *entities.Route) error
	GetRouteByPath(ctx context.Context, path string) (*entities.Route, error)
	GetRouteByPathAndMethod(ctx context.Context, path string, method entities.HttpMethod) (*entities.Route, error)
	GetRouteByID(ctx context.Context, id uuid.UUID) (*entities.Route, error)
	UpdateRoute(ctx context.Context, route *entities.Route) error
	DeleteRoute(ctx context.Context, id uuid.UUID) error
}
