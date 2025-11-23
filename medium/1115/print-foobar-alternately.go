package leetcode

type FooBar struct {
	n       int
	fooLock chan any
	barLock chan any
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:       n,
		fooLock: make(chan any),
		barLock: make(chan any, 1),
	}

	fb.barLock <- 1

	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barLock
		printFoo()
		fb.fooLock <- 1
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooLock
		printBar()
		fb.barLock <- 1
	}
}
