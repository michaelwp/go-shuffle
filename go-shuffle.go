package go_shuffle

/*
author: michael.wenceslaus@gmail.com
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

// Shuffle array list in 'n' times
// example :
// var list = goshuffle.List{1,"a",3,"z",5}
// fmt.Println(list.Shuffle())
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

// Shuffle array 1 times
// example :
// var list = goshuffle.List{1,"a",3,"z",5}
//	var listTimes = goshuffle.ListTimes{
//		List:  list,
//		Times: 5,
//	}
//
//	fmt.Println(listTimes.Shuffle())
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
