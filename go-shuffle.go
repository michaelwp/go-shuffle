package go_shuffle

import (
	"math/rand"
	"time"
)

type List []interface{}
type Times int

// Start recursive
func (l List) Start(t Times) []interface{} {
	var result []interface{}

	for i := 0; i < int(t); i++ {
		var r []interface{}
		res := l.shuffle()

		for _, x := range res {
			r = append(r, x)
		}

		result = append(result, r)
	}

	return result
}

func (l List) shuffle() List {
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