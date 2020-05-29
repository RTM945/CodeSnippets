package syncbychan

// type Mutex struct {
// 	c      chan struct{}
// 	locked bool
// }

// func NewMutex() *Mutex {
// 	return &Mutex{
// 		c:      make(chan struct{}, 1),
// 		locked: false,
// 	}
// }

// func (m *Mutex) Lock() {
// 	if m.locked {
// 		panic("lock of locked mutex")
// 	} else {
// 		m.c <- struct{}{}
// 		m.locked = true
// 	}
// }

// func (m *Mutex) Unlock() {
// 	if m.locked {
// 		<-m.c
// 		m.locked = false
// 	} else {
// 		panic("unlock of unlocked mutex")
// 	}
// }

type Mutex chan struct{}

func NewMutex() Mutex {
	return make(chan struct{}, 1)
}

func (m Mutex) Lock() {
	m <- struct{}{}
}

func (m Mutex) Unlock() {
	<-m
}
