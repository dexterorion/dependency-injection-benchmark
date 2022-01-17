package shared

func NewFakeService() (FakeService, error) { // dingo força fazermos isso
	return &fakeServiceImpl{}, nil
}

type FakeService interface{}

type fakeServiceImpl struct{}
