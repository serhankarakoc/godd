# Golang debug dump and die (dd)
Golang debug dump and die (dd)

# Install
```bash
go get github.com/serhankarakoc/go-dd
```

# Usage 
Dump Struct:
```go
package main

import (
	"github.com/serhankarakoc/go-dd"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	jack := Person{
		Name: "jack",
		Age:  32,
	}
	dump.DD(jack)
}

// result

#Type: main.Person
#Underlying Type: struct
#Value: {jack 32}
```


Dump Strings:
```go
package main

import (
	"github.com/serhankarakoc/go-dd"
)

func main() {
	t := "this is a test string"
	dump.DD(t)
}

// result

#Type: string
#Value: this is a test string
```

Dump Numeric:
```go
package main

import (
	"github.com/serhankarakoc/go-dd"
)

func main() {
	t := 123.43
	dump.DD(t)
}

// result

#Type: float64
#Value: 123.43
```

Dump Channel:
```go
package main

import (
	"github.com/serhankarakoc/go-dd"
)

func main() {
	var c chan int
	dump.DD(c)
}

// result

#Type: chan int
#Underlying Type: chan
#Value: <nil>
```

Dump Interface:
```go
package main

import (
	"github.com/serhankarakoc/go-dd"
)

func main() {
	type Getter interface {
		get()
	}
	var g Getter
	dump.DD(g)
}

// result

#Type: interface
```
