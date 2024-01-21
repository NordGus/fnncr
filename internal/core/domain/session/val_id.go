package session

const (
	idByteSize = 64
)

type ID [idByteSize]byte

func NewID(id [idByteSize]byte) (ID, error) {
	return id, nil
}

func (id ID) String() string {
	panic("unimplemented")
}
