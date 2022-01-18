package services

import (
	"fmt"

	"github.com/dexterorion/dependency-injection-benchmark/shared/storage"
)

func NewFakeService(store storage.FakeStorage) (FakeService, error) { // dingo força fazermos isso
	return &fakeServiceImpl{
		storage: store,
	}, nil
}

type FakeService interface {
	String() string
}

type fakeServiceImpl struct {
	storage storage.FakeStorage
}

func (fs *fakeServiceImpl) String() string {
	return fmt.Sprint(fs.storage)
}

// ---- fake struct

func NewFakeServiceStruct(store storage.FakeStorage) (*FakeServiceStruct, error) { // dingo força fazermos isso
	return &FakeServiceStruct{
		storage: store,
	}, nil
}

type FakeServiceStruct struct {
	storage storage.FakeStorage
}

func (fss *FakeServiceStruct) String() string {
	// return "struct"
	return fmt.Sprint(fss.storage)
}
