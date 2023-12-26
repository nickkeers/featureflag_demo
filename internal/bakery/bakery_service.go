package bakery

import (
	"context"
	"github.com/open-feature/go-sdk/openfeature"
)

type Service struct {
	flagClient *openfeature.Client
}

type Cake struct {
	Flour string
}

func NewBakeryService(client *openfeature.Client) *Service {
	return &Service{flagClient: client}
}

func (b *Service) BakeCake(ctx context.Context) *Cake {
	flagEnabled, _ := b.flagClient.BooleanValue(
		ctx, "glutenFree", true, openfeature.EvaluationContext{},
	)

	if flagEnabled {
		return &Cake{Flour: "almond"}
	} else {
		return &Cake{Flour: "normal"}
	}
}
