package leetcode

type ATM struct {
	set     [5]int
	balance [5]int
}

func Constructor() ATM {
	return ATM{
		balance: [5]int{},
		set:     [5]int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, v := range banknotesCount {
		this.balance[i] += v
	}
}

func (this *ATM) Withdraw(amount int) []int {
	var curr int = amount
	var result [5]int
	for i := len(this.balance) - 1; i >= 0; i-- {
		if this.set[i] > curr || this.balance[i] == 0 {
			continue
		}

		count := curr / this.set[i]
		if this.balance[i] >= count {
			result[i] = count
			curr -= (this.set[i] * count)
		} else {
			result[i] = this.balance[i]
			curr -= (this.set[i] * this.balance[i])
		}
	}

	if curr != 0 {
		return []int{-1}
	}

	for i, v := range result {
		this.balance[i] -= v
	}

	return result[:]
}
