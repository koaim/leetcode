package leetcode

type Foo struct {
	secondAllow chan any
	thirdAllow  chan any
}

/*
The same instance of Foo will be passed to three different threads.
Thread A will call first(), thread B will call second(), and thread C will call third().

Design a mechanism and modify the program to ensure that second() is executed after first(), and third() is executed after second().
*/
func NewFoo() *Foo {
	return &Foo{
		secondAllow: make(chan any),
		thirdAllow:  make(chan any),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	f.secondAllow <- 1
}

func (f *Foo) Second(printSecond func()) {
	<-f.secondAllow
	printSecond()
	f.thirdAllow <- 1
}

func (f *Foo) Third(printThird func()) {
	<-f.thirdAllow
	printThird()
}
