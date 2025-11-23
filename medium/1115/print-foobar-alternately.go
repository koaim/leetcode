package leetcode

type FooBar struct {
	n       int
	fooLock chan any
	barLock chan any
}

/*
Suppose you are given the following code:

class FooBar {
  public void foo() {
    for (int i = 0; i < n; i++) {
      print("foo");
    }
  }

  public void bar() {
    for (int i = 0; i < n; i++) {
      print("bar");
    }
  }
}

The same instance of FooBar will be passed to two different threads:

	thread A will call foo(), while
	thread B will call bar().

Modify the given program to output "foobar" n times.
*/
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
