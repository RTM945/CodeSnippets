package syncbychan

type genaration struct {
	wait chan struct{}
	n    int
}

func newGenaration() genaration {
	return genaration{wait: make(chan struct{})}
}

func (g genaration) end() {
	close(g.wait)
}

type WaitGroup chan genaration

func NewWaitGroup() WaitGroup {
	wg := make(WaitGroup, 1)
	g := newGenaration()
	g.end()
	wg <- g
	return wg
}

func (wg WaitGroup) Add(delta int) {
	g := <-wg
	if g.n == 0 {
		g = newGenaration()
	}
	g.n += delta
	if g.n < 0 {
		panic("negative WaitGroup count")
	}
	if g.n == 0 {
		g.end()
	}
	wg <- g
}

func (wg WaitGroup) Done() {
	wg.Add(-1)
}

func (wg WaitGroup) Wait() {
	g := <-wg
	wait := g.wait
	wg <- g
	<-wait
}
