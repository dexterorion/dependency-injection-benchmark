package storage

type fakeLazy struct{}

func NewFakeLazy() (FakeStorage, error) {
	// access database on each call
	return &fakeLazy{}, nil
}

func (fm *fakeLazy) Save() {}

func (fm *fakeLazy) Delete() {}

func (fm *fakeLazy) List() {}

func (fr *fakeLazy) String() string {
	return "lazy"
}
