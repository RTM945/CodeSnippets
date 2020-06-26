package syncbychan

type Mutex chan struct{}

func NewMutex() Mutex {
	return make(Mutex, 1)
}

func (m Mutex) Lock() {
	m <- struct{}{}
}

func (m Mutex) Unlock() {
	<-m
}

func (m Mutex) TryLock() bool {
	select {
	case m <- struct{}{}:
		return true
	default:
		return false
	}
}
