package dic

import (
	"errors"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	services "github.com/dexterorion/dependency-injection-benchmark/shared/services"
	storage "github.com/dexterorion/dependency-injection-benchmark/shared/storage"
)

func getDiDefs(provider dingo.Provider) []di.Def {
	return []di.Def{
		{
			Name:  "fakeservice",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("fakeservice")
				if err != nil {
					var eo services.FakeService
					return eo, err
				}
				pi0, err := ctn.SafeGet("fakestorageinmem")
				if err != nil {
					var eo services.FakeService
					return eo, err
				}
				p0, ok := pi0.(storage.FakeStorage)
				if !ok {
					var eo services.FakeService
					return eo, errors.New("could not cast parameter 0 to storage.FakeStorage")
				}
				b, ok := d.Build.(func(storage.FakeStorage) (services.FakeService, error))
				if !ok {
					var eo services.FakeService
					return eo, errors.New("could not cast build function to func(storage.FakeStorage) (services.FakeService, error)")
				}
				return b(p0)
			},
			Unshared: false,
		},
		{
			Name:  "fakestorageinmem",
			Scope: "",
			Build: func(ctn di.Container) (interface{}, error) {
				d, err := provider.Get("fakestorageinmem")
				if err != nil {
					var eo storage.FakeStorage
					return eo, err
				}
				b, ok := d.Build.(func() (storage.FakeStorage, error))
				if !ok {
					var eo storage.FakeStorage
					return eo, errors.New("could not cast build function to func() (storage.FakeStorage, error)")
				}
				return b()
			},
			Unshared: false,
		},
	}
}
