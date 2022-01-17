package provider

import (
	"context"

	"github.com/dexterorion/dependency-injection-benchmark/shared"
	"github.com/sarulabs/dingo/v4"
)

type Provider struct {
	dingo.BaseProvider
}

func (p *Provider) Load() error {
	if err := p.AddDefSlice([]dingo.Def{{Name: "fakeservice", Build: shared.NewFakeService}}); err != nil {
		return err
	}
	if err := p.AddDefSlice([]dingo.Def{{Name: "dsclient", Build: shared.GetDatastoreClient, Params: dingo.NewFuncParams(context.Background(), shared.DIBenchmarkProjectId)}}); err != nil {
		return err
	}
	return nil
}
