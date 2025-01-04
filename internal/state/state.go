package state

type State interface {
	Append(mark string) error
}

type Store interface {
	Append(mark string) error
}

type state struct {
	store Store
}
