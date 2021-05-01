# go-shuffle v1.2.2

Ini adalah package bahasa pemrograman go sederhana untuk mengacak susunan list array.

### func Shuffle()

```go
func (l List) Shuffle() List
```
```go
func (l ListTimes) Shuffle() List
```

### Instalasi

`go get -u github.com/michaelwp/go-shuffle`

### Contoh penggunaan

```go 
package main

import (
	"fmt"
	goshuffle "github.com/michaelwp/go-shuffle"
)

func main()  {
	var list = goshuffle.List{1,"a",3,"z",5}
	var listTimes = goshuffle.ListTimes{
		List:  list,
		Times: 5,
	}

	fmt.Println(list.Shuffle())
	fmt.Println(listTimes.Shuffle())
}
```

#### Hasil

```text
[a 3 z 5 1]
[[5 1 3 z a] [3 5 a 1 z] [z 3 5 a 1] [a 1 3 5 z] [3 5 1 z a]]
```