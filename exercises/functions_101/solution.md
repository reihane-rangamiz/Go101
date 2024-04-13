# Solution

## 1.

```go
package main

import "fmt"

func sumTwoInt(x, y int) int {
	z := x + y
	return z
}

func main() {

	fmt.Println("The result is: ", sumTwoInt(1, 2))
}
```

## 2. 

When a function has a return value it must specify it. the correct form would be :

```go
func add(x, y int) int{
    return x + y
}
```