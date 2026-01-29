package mode

import "sync/atomic"

type Mode string

const (
	OK        Mode = "ok"
	Error     Mode = "error"
	RareError Mode = "rare_error"
	Slow      Mode = "slow"
	Flaky     Mode = "flaky"
)

type Store struct {
	mode atomic.Value
}

func NewModeStore() *Store {
	s := &Store{}
	s.mode.Store(OK)
	return s
}

func (s *Store) Set(mode Mode) {
	s.mode.Store(mode)
}

func (s *Store) Get() Mode {
	return s.mode.Load().(Mode)
}
