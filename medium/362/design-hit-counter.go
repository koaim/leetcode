package leetcode

/*
Design a hit counter which counts the number of hits received in the past 5 minutes (i.e., the past 300 seconds).
Your system should accept a timestamp parameter (in seconds granularity), and you may assume that calls are being made to the system in chronological order (i.e., timestamp is monotonically increasing). Several hits may arrive roughly at the same time.

Implement the HitCounter class:
	- HitCounter() Initializes the object of the hit counter system.
    - void hit(int timestamp) Records a hit that happened at timestamp (in seconds). Several hits may happen at the same timestamp.
    - int getHits(int timestamp) Returns the number of hits in the past 5 minutes from timestamp (i.e., the past 300 seconds).
*/
type HitCounter struct {
	calls []int
}

func Constructor() HitCounter {
	return HitCounter{
		calls: []int{},
	}
}

func (this *HitCounter) Hit(timestamp int) {
	this.calls = append(this.calls, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	for len(this.calls) > 0 && this.calls[0] < timestamp-299 {
		this.calls = this.calls[1:]
	}

	return len(this.calls)
}
