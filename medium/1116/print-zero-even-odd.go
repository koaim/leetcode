package leetcode

type ZeroEvenOdd struct {
	n       int
	canZero chan any
	canEven chan any
	canOdd  chan any
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:       n,
		canZero: make(chan any, 1),
		canEven: make(chan any),
		canOdd:  make(chan any),
	}

	zeo.canZero <- 1
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for v := 0; v < z.n; v++ {
		<-z.canZero

		if v != z.n {
			printNumber(0)
		}

		if (v+1)%2 == 0 {
			z.canEven <- 1
		} else {
			z.canOdd <- 1
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for v := 1; v <= z.n; v++ {
		if (v % 2) == 0 {
			<-z.canEven
			printNumber(v)
			z.canZero <- 1
		}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for v := 1; v <= z.n; v++ {
		if (v % 2) != 0 {
			<-z.canOdd
			printNumber(v)
			z.canZero <- 1
		}
	}
}
