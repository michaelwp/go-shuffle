# go-shuffle v1.0.0

Ini adalah package bahasa pemrograman go sederhana untuk mengacak susunan list array.

### func Start()

```go
func (l List) Start(t Times) List
```

### func Shuffle()

```go
func (l List) Shuffle() List
```

### Instalasi

`go get -u github.com/michaelwp/go-shuffle`

### Contoh penggunaan

```go 
package main

import (
	"fmt"
	go_shuffle "github.com/michaelwp/go-shuffle"
)

func main()  {
	var list = go_shuffle.List{1,"a",3,"z",5}

	fmt.Println(list.Shuffle())
	fmt.Println(list.Start(5))
}
```

#### Hasil

```text
[3 z a 5 1]
[[a 1 z 3 5] [5 a 1 z 3] [1 5 3 a z] [5 z 1 3 a] [z 1 3 a 5]]
```