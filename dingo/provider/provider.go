package provider

import (
	"github.com/dexterorion/dependency-injection-benchmark/shared/services"
	"github.com/dexterorion/dependency-injection-benchmark/shared/storage"
	"github.com/sarulabs/dingo/v4"
)

type Provider struct {
	dingo.BaseProvider
}

func (p *Provider) Load() error {
	if err := p.AddDefSlice([]dingo.Def{{Name: "fakeservice", Build: services.NewFakeService}}); err != nil {
		return err
	}
	if err := p.AddDefSlice([]dingo.Def{{Name: "fakestorageinmem", Build: storage.NewFakeInMem}}); err != nil {
		return err
	}
	// if err := p.AddDefSlice([]dingo.Def{{Name: "fakestoragelazy", Build: storage.NewFakeLazy}}); err != nil {
	// 	return err
	// } // cannot do that
	return nil
}
