# Loops 102

## Solution

1. 
```go
package main

import "fmt"

func main(){

    var colors = [3]string{"red","green","blue"}

    for _,color := range colors{
        fmt.Println(color)
    }
}

```

2.
```go
package main

import "fmt"

func main(){

    var x = 1

    for x < 101 {
        x += 3
    }

    fmt.Printf("The final value is: %v \n",x)
}

```