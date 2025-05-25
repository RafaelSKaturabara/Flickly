package core

import "context"

type Service interface {
	AbleToRun(ctx context.Context, entity Entity) bool
	Run(ctx context.Context, entity Entity) error
}
