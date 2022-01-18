package storage

type FakeStorage interface {
	Save()
	Delete()
	List()

	String() string
}
