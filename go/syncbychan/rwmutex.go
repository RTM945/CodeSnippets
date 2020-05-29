package syncbychan

type RWMutex struct {
	w chan struct{}
	r chan int
}

func NewRWMutex() *RWMutex {
	return &RWMutex{
		w: make(chan struct{}, 1),
		r: make(chan int, 1),
	}
}

func (m *RWMutex) Lock() {
	m.w <- struct{}{}
}

func (m *RWMutex) UnLock() {
	<-m.w
}

// int readerCount = 0;
// int R = 1;
// int W = 1;
// read() {
// 	while(1) {
// 		P(R);
// 		readerCount++;
// 		if (readerCount > 0) P(W);
// 		V(R);
// 		//read ...
// 		P(R);
// 		readerCount--;
// 		if (readerCount == 0) V(W);
// 		V(R);
// 	}
// }
// write() {
// 	while(1) {
// 		P(W)
// 		// writer
// 		V(W)
// 	}
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
	} else {
		m.r <- rs
	}
}
