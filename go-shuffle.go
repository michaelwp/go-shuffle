package go_shuffle

/**
v.1.3.0
*/

import (
	"math/rand"
	"time"
)

type List []interface{}
type ListTimes struct {
	List  List
	Times int
}

func (l ListTimes) Shuffle() List {
	var result []interface{}

	for i := 0; i < l.Times; i++ {
		var r []interface{}
		res := l.List.Shuffle()

		for _, x := range res {
			r = append(r, x)
		}

		result = append(result, r)
	}

	return result
}

func (l List) Shuffle() List {
	var length = len(l) - 1

	// set new random seed
	rand.Seed(time.Now().UnixNano())

	// Fisher–Yates shuffle
	for i := length; i > 0; i-- {
		r := rand.Intn(i)
		l[i], l[r] = l[r], l[i]
	}

	return l
}
