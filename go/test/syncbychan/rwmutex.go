package syncbychan

type RWMutex struct {
	r chan int
	w chan struct{}
}

func NewRWMutex() *RWMutex {
	return &RWMutex{
		w: make(chan struct{}, 1),
		r: make(chan int, 1),
	}
}

// int reader = 0;
// int mutex = 1;
// int write = 1;

// //writer
// while(1) {
// 	P(write)
// 	// writer
// 	V(write)
// }

// //reader
// while(1) {
// 	P(mutex)
// 	reader++
// 	if(reader > 0) {
// 		P(write)
// 	}
// 	V(mutex)
// 	//read
// 	P(mutex)
// 	reader--
// 	if(reader = 0) {
// 		V(write)
// 	}
// 	V(mutex)
// }

func (m *RWMutex) RLock() {
	var rs int
	select {
	case m.w <- struct{}{}:
	case rs = <-m.r:
	}
	rs++
	m.r <- rs
}

func (m *RWMutex) RUnlock() {
	rs := <-m.r
	rs--
	if rs == 0 {
		<-m.w
		return
		// 不必写回 之后再有加锁 就是从0开始的新rs
	}
	m.r <- rs
}

func (m *RWMutex) Lock() {
	m.w <- struct{}{}
}

func (m *RWMutex) Unlock() {
	<-m.w
}

func (m *RWMutex) TryLock() bool {
	select {
	case m.w <- struct{}{}:
		return true
	default:
		return false
	}
}

func (m *RWMutex) TryRLock() bool {
	var rs int
	select {
	case m.w <- struct{}{}:
	case rs = <-m.r:
	default:
		return false
	}
	rs++
	m.r <- rs
	return true
}
