# Golang debug dump and die (dd)

# Install
```bash
go get github.com/serhankarakoc/godd
```

# Usage 
Dump Struct:
```go
package main

import (
	"github.com/serhankarakoc/godd"
)

func main() {
	type Animal struct {
		Name string
		Age  int
	}

	leo := Animal {
		Name: "leo",
		Age:  5,
	}
	godd.DD(leo)
}

// result
Type: main.Animal
Underlying Type: struct
Value: {leo 5}
```


Dump Strings:
```go
package main

import (
	"github.com/serhankarakoc/godd"
)

func main() {
	t := "test string"
	godd.DD(t)
}

// result
Type: string
Value: test string
```

Dump Numeric:
```go
package main

import (
	"github.com/serhankarakoc/godd"
)

func main() {
	t := 123.45
	godd.DD(t)
}

// result
Type: float64
Value: 123.45
```

Dump Channel:
```go
package main

import (
	"github.com/serhankarakoc/godd"
)

func main() {
	var c chan int
	godd.DD(c)
}

// result
Type: chan int
Underlying Type: chan
Value: <nil>
```

Dump Interface:
```go
package main

import (
	"github.com/serhankarakoc/godd"
)

func main() {
	type Getter interface {
		get()
	}
	var g Getter
	godd.DD(g)
}

// result
Type: interface
```
