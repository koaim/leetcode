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
	for i := 1; i <= z.n; i++ {
		<-z.canZero
		printNumber(0)

		if i%2 == 0 {
			z.canEven <- 1
		} else {
			z.canOdd <- 1
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		if i%2 == 0 {
			<-z.canEven
			printNumber(i)
			z.canZero <- 1
		}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= z.n; i++ {
		if i%2 != 0 {
			<-z.canOdd
			printNumber(i)
			z.canZero <- 1
		}
	}
}
