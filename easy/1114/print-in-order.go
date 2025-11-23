package leetcode

type Foo struct {
	first  chan any
	second chan any
}

/*
The same instance of Foo will be passed to three different threads.
Thread A will call first(), thread B will call second(), and thread C will call third().

Design a mechanism and modify the program to ensure that second() is executed after first(), and third() is executed after second().
*/
func NewFoo() *Foo {
	return &Foo{
		first:  make(chan any),
		second: make(chan any),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	f.first <- 1
}

func (f *Foo) Second(printSecond func()) {
	<-f.first
	printSecond()
	f.second <- 1
}

func (f *Foo) Third(printThird func()) {
	<-f.second
	printThird()
}
