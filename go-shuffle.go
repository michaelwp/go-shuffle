package go_shuffle

/**
v.1.2.0
*/

import (
	"math/rand"
	"time"
)

type List []interface{}
type Times int

type Shuffle interface {
	Start() List
	Times(t Times) List
}

func (l List) Times(t Times) List {
	var result []interface{}

	for i := 0; i < int(t); i++ {
		var r []interface{}
		res := l.Start()

		for _, x := range res {
			r = append(r, x)
		}

		result = append(result, r)
	}

	return result
}

func (l List) Start() List {
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
