package bakery

import (
	"context"
	"github.com/open-feature/go-sdk/openfeature"
	. "github.com/open-feature/go-sdk/openfeature/memprovider"
	"strconv"
	"testing"
)

func buildMemoryProvider(flags map[string]bool) openfeature.FeatureProvider {
	flagMap := make(map[string]InMemoryFlag)

	for key, val := range flags {
		flagMap[key] = InMemoryFlag{
			Key:            key,
			State:          Enabled,
			DefaultVariant: strconv.FormatBool(val),
			Variants: map[string]interface{}{
				"true":  true,
				"false": false,
			},
			ContextEvaluator: nil,
		}
	}

	return NewInMemoryProvider(flagMap)
}

func TestService_BakeCake(t *testing.T) {
	t.Run("when flag is true", func(t *testing.T) {
		openfeature.SetProvider(buildMemoryProvider(map[string]bool{
			"glutenFree": true,
		}))

		client := openfeature.NewClient("flaaags")
		bakeryService := NewBakeryService(client)

		cake := bakeryService.BakeCake(context.Background())

		if cake.Flour != "almond" {
			t.Errorf("expected almond flour, got %s", cake.Flour)
		}
	})

	t.Run("when flag is false", func(t *testing.T) {
		openfeature.SetProvider(buildMemoryProvider(map[string]bool{
			"glutenFree": false,
		}))

		client := openfeature.NewClient("flaaags")
		bakeryService := NewBakeryService(client)

		cake := bakeryService.BakeCake(context.Background())

		if cake.Flour != "normal" {
			t.Errorf("expected normal flour, got %s", cake.Flour)
		}
	})
}
