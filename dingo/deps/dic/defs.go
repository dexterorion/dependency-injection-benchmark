package dic

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	context "context"

	datastore "cloud.google.com/go/datastore"
	shared "github.com/dexterorion/dependency-injection-benchmark/shared"
)

func getDiDefs(provider dingo.Provider) []di.Def {
	return []di.Def{
		{
			Name:  "dsclient",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("dsclient")
				if err != nil {
					var eo *datastore.Client
					return eo, err
				}
				pi0, ok := d.Params["0"]
				if !ok {
					var eo *datastore.Client
					return eo, errors.New("could not find parameter 0")
				}
				p0, ok := pi0.(context.Context)
				if !ok {
					var eo *datastore.Client
					return eo, errors.New("could not cast parameter 0 to context.Context")
				}
				pi1, ok := d.Params["1"]
				if !ok {
					var eo *datastore.Client
					return eo, errors.New("could not find parameter 1")
				}
				p1, ok := pi1.(string)
				if !ok {
					var eo *datastore.Client
					return eo, errors.New("could not cast parameter 1 to string")
				}
				b, ok := d.Build.(func(context.Context, string) (*datastore.Client, error))
				if !ok {
					var eo *datastore.Client
					return eo, errors.New("could not cast build function to func(context.Context, string) (*datastore.Client, error)")
				}
				return b(p0, p1)
			},
			Unshared: false,
		},
		{
			Name:  "fakeservice",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("fakeservice")
				if err != nil {
					var eo shared.FakeService
					return eo, err
				}
				b, ok := d.Build.(func() (shared.FakeService, error))
				if !ok {
					var eo shared.FakeService
					return eo, errors.New("could not cast build function to func() (shared.FakeService, error)")
				}
				return b()
			},
			Unshared: false,
		},
	}
}
