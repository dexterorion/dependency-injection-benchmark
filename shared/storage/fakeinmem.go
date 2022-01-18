package storage

type fakeInMem struct {
	mem map[string]interface{}
}

func NewFakeInMem() (FakeStorage, error) {
	// keeps listening to messages and updates memory
	store := &fakeInMem{}

	go store.listen()

	return &fakeInMem{}, nil
}

func (fr *fakeInMem) Save() {}

func (fr *fakeInMem) Delete() {}

func (fr *fakeInMem) List() {}

func (fr *fakeInMem) listen() {
	// listen to messages for InMem (or pubsub)
	fr.mem = make(map[string]interface{})
}

func (fr *fakeInMem) String() string {
	return "inmem"
}
