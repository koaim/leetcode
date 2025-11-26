package leetcode

type H2O struct {
	hSem  chan any
	oSem  chan any
	mutex chan any
	count int
}

func NewH2O() *H2O {
	return &H2O{
		hSem:  make(chan any, 2),
		oSem:  make(chan any, 1),
		mutex: make(chan any, 1),
	}
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.hSem <- 1
	releaseHydrogen()

	h.mutex <- 1
	h.count++
	if h.count == 2 {
		h.oSem <- 1
	}
	<-h.mutex
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	<-h.oSem
	releaseOxygen()

	h.mutex <- 1
	h.count = 0
	<-h.mutex

	<-h.hSem
	<-h.hSem
}
