package sample

type ID string

func (id ID) String() string {
	return string(id)
}
